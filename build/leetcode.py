import time
import json
import os
import sys
import requests
from requests.adapters import HTTPAdapter
requests.packages.urllib3.disable_warnings()


def login(EMAIL, PASSWORD):
    session = requests.Session()
    session.mount('http://', HTTPAdapter(max_retries=6))
    session.mount('https://', HTTPAdapter(max_retries=6))
    session.encoding = "utf-8"

    login_data = {
        'login': EMAIL,
        'password': PASSWORD
    }

    sign_in_url = 'https://leetcode.cn/accounts/login/'
    headers = {'User-Agent': "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.122 Safari/537.36",
               'Connection': 'keep-alive', 'Referer': sign_in_url, "origin": "https://leetcode.cn/"}

    session.post(sign_in_url, headers=headers, data=login_data,
                 timeout=10, allow_redirects=False)
    is_login = session.cookies.get('LEETCODE_SESSION') != None
    if is_login:
        return session
    else:
        raise Exception("login error")


def get_submission_list(slug, session):
    url = 'https://leetcode.cn/graphql/'

    payload = json.dumps({
        "operationName": "submissions",
        "variables": {
            "offset": 0,
            "limit": 40,
            "lastKey": "null",
            "questionSlug": slug
        },
        "query": "query submissions($offset: Int!, $limit: Int!, $lastKey: String, $questionSlug: String!) {\n  submissionList(offset: $offset, limit: $limit, lastKey: $lastKey, questionSlug: $questionSlug) {\n    lastKey\n    hasNext\n    submissions {\n      id\n      statusDisplay\n      lang\n      runtime\n      timestamp\n      url\n      isPending\n      memory\n      __typename\n    }\n    __typename\n  }\n}\n"
    })

    headers = {"content-type": "application/json", "origin": "https://leetcode.cn", "referer": "https://leetcode.cn/progress/",
               "user-agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.122 Safari/537.36"}

    r = session.post(url, data=payload, headers=headers, verify=False)
    response_data = json.loads(r.text)

    return response_data


def get_accepted_problems(session):
    url = 'https://leetcode.cn/graphql/'

    payload = json.dumps({
        "operationName": "userProfileQuestions",
        "variables": {
            "status": "ACCEPTED",
            "skip": 0,
            "first": 200000,  # 一次返回的数据量
            "sortField": "LAST_SUBMITTED_AT",
            "sortOrder": "DESCENDING",
            "difficulty": [
            ]
        },
        "query": "query userProfileQuestions($status: StatusFilterEnum!, $skip: Int!, $first: Int!, $sortField: SortFieldEnum!, $sortOrder: SortingOrderEnum!, $keyword: String, $difficulty: [DifficultyEnum!]) {\n  userProfileQuestions(status: $status, skip: $skip, first: $first, sortField: $sortField, sortOrder: $sortOrder, keyword: $keyword, difficulty: $difficulty) {\n    totalNum\n    questions {\n      translatedTitle\n      frontendId\n      titleSlug\n      title\n      difficulty\n      lastSubmittedAt\n      numSubmitted\n      lastSubmissionSrc {\n        sourceType\n        ... on SubmissionSrcLeetbookNode {\n          slug\n          title\n          pageId\n          __typename\n        }\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n}\n"
    })

    headers = {"content-type": "application/json", "origin": "https://leetcode.cn", "referer": "https://leetcode.cn/progress/",
               "user-agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.122 Safari/537.36"}

    r = session.post(url, data=payload, headers=headers, verify=False)
    response_data = json.loads(r.text)

    return response_data


def list_local_solution():
    dir_name = 'leetcode'
    local_su = {}
    for dirpath, dirname, filenames in os.walk(dir_name):
        if filenames:
            for file in filenames:
                if not file.endswith('_test.go'):
                    id = file.split('.')[0]
                    local_su[id] = [file, dirpath + '/' + file]
    print(local_su)
    return local_su


def generate_markdown_text(response_data, session):
    markdown_text = "# leetcode-golang\n"
    markdown_text += "\n"
    markdown_text += "```go\n"
    markdown_text += "package main\n"
    markdown_text += "\n"
    markdown_text += "import \"fmt\"\n"
    markdown_text += "\n"
    markdown_text += "func main() {\n"
    markdown_text += "  fmt.Println(\"Hello, leetcode\")\n"
    markdown_text += "}\n"
    markdown_text += "```\n"
    markdown_text += "\n"
    markdown_text += "# 提交总览"
    markdown_text += "\n"
    markdown_text += "| 编号 | 题目 | 解法 | 题目难度 | 提交次数| 重刷次数 |\n| --- | --- | --- | --- | --- | --- |\n"
    table_body = []

    response_data = response_data['data']['userProfileQuestions']['questions']

    localSu = list_local_solution()
    
    for index, sub_data in enumerate(response_data):

        print('{}/{}'.format(index + 1, len(response_data)))

        lastSubmittedAt = time.strftime(
            "%Y-%m-%d %H:%M", time.localtime(sub_data['lastSubmittedAt']))
        translatedTitle = "#{} {}".format(
            sub_data['frontendId'], sub_data['translatedTitle'])
        frontendId = sub_data['frontendId']
        difficulty = sub_data['difficulty']
        titleSlug = sub_data['titleSlug']
        numSubmitted = sub_data['numSubmitted']
        numSubmitted = str(numSubmitted)
        url = "https://leetcode.cn/problems/{}".format(
            sub_data['titleSlug'])
        
        localFile = ''
        localLink = ''
        if frontendId in localSu.keys():
            localFile = localSu[frontendId][0]
            localLink = localSu[frontendId][1]
        localText = ''
        if localFile != '':
            localText = '[' + localFile + ']' + '(' + localLink + ')'

        submission_dict = get_submission_list(titleSlug, session)
        submission_list = submission_dict['data']['submissionList']['submissions']
        submission_accepted_dict = {}

        for submission in submission_list:
            status = submission['statusDisplay']
            if (status == 'Accepted'):
                submission_time = time.strftime(
                    "%Y-%m-%d", time.localtime(int(submission['timestamp'])))
                if submission_time in submission_accepted_dict.keys():
                    submission_accepted_dict[submission_time] += 1
                else:
                    submission_accepted_dict[submission_time] = 1

        count = len(submission_accepted_dict)
        if count > 1:
            count = "**" + str(count) + "**"
        else:
            count = str(count)
        table_body.append("|[%s](%s)|%s|%s|%s|%s|" % (
            translatedTitle, url, localText, difficulty, numSubmitted, count))

    table_body = sorted(table_body, key=lambda ele: int(
        ele.split('#')[1].split(' ')[0]))
    for idx, ele in enumerate(table_body):
        markdown_text += "|" + str(idx+1) + ele + '\n'
    markdown_text += "\n"
    markdown_text += "# 文件结构"
    markdown_text += "\n"
    markdown_text += "![Visualization of this repo](./diagram.svg)"
    markdown_text += "\n"

    with open('Algo.md', 'r') as f:
        markdown_text += f.read()

    return markdown_text


if __name__ == '__main__':
    session = login(sys.argv[1], sys.argv[2])
    response_data = get_accepted_problems(session)
    markdown_text = generate_markdown_text(response_data, session)
    with open("README.md", "w") as f:
        f.write(markdown_text)