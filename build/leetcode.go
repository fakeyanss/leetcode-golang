package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

// 题目信息结构体
type LeetCodeProblem struct {
	Num        int    // 题目序号
	Title      string // 题目名称
	URL        string // 题目链接
	Difficulty string // 难度（简单/中等/困难）
	Solution   string // 解法描述
}

const (
	// 配置项
	codeDir          = "leetcode"  // 刷题代码目录
	readmePath       = "README.md" // README路径
	tableStartMarker = "<!-- LEETCODE_TABLE_START -->"
	tableEndMarker   = "<!-- LEETCODE_TABLE_END -->"
)

// 遍历目录解析所有题目文件
func parseAllProblems() ([]LeetCodeProblem, error) {
	var problems []LeetCodeProblem

	// 遍历指定目录下的所有.go文件
	err := filepath.Walk(codeDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// 跳过目录、测试文件
		if info.IsDir() || strings.HasSuffix(info.Name(), "_test.go") {
			return nil
		}
		// 只处理.go文件
		if filepath.Ext(info.Name()) != ".go" {
			return nil
		}

		// 解析单个文件
		problem, err := parseProblemFile(path)
		if err != nil {
			fmt.Printf("解析文件失败 %s: %v\n", path, err)
			return nil // 跳过解析失败的文件，继续处理其他
		}
		if problem.Num > 0 { // 有效题目才加入
			problems = append(problems, problem)
		}
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("遍历目录失败: %w", err)
	}

	// 按题目序号排序
	sort.Slice(problems, func(i, j int) bool {
		return problems[i].Num < problems[j].Num
	})

	return problems, nil
}

// 解析单个题目文件
func parseProblemFile(filePath string) (LeetCodeProblem, error) {
	var problem LeetCodeProblem
	fileName := filepath.Base(filePath)

	// 1. 从文件名提取题目序号（匹配 lc151/151.xxx.go 格式）
	numRegex := regexp.MustCompile(`lc(\d+)|(\d+)\.`)
	numMatches := numRegex.FindStringSubmatch(fileName)
	if numMatches != nil {
		var numStr string
		if numMatches[1] != "" {
			numStr = numMatches[1]
		} else if numMatches[2] != "" {
			numStr = numMatches[2]
		}
		_, _ = fmt.Sscanf(numStr, "%d", &problem.Num)
	}
	if problem.Num == 0 {
		return problem, fmt.Errorf("未提取到题目序号")
	}

	// 2. 读取文件内容解析注释
	file, err := os.Open(filePath)
	if err != nil {
		return problem, fmt.Errorf("打开文件失败: %w", err)
	}
	defer file.Close()

	var content strings.Builder
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		content.WriteString(scanner.Text() + "\n")
	}
	if err := scanner.Err(); err != nil {
		return problem, fmt.Errorf("读取文件内容失败: %w", err)
	}
	fileContent := content.String()

	// 3. 提取题目标题
	// 优先匹配 // [151] 反转字符串中的单词 格式
	titleRegex1 := regexp.MustCompile(`\[(\d+)\]\s*([^\n]+)`)
	titleMatches1 := titleRegex1.FindStringSubmatch(fileContent)
	if titleMatches1 != nil && len(titleMatches1) >= 3 {
		problem.Title = strings.TrimSpace(titleMatches1[2])
	} else {
		// 备用：从文件名提取（151.反转字符串中的单词.go → 反转字符串中的单词）
		titleRegex2 := regexp.MustCompile(`\d+\.([^.]+)\.go`)
		titleMatches2 := titleRegex2.FindStringSubmatch(fileName)
		if titleMatches2 != nil && len(titleMatches2) >= 2 {
			problem.Title = strings.TrimSpace(titleMatches2[1])
		} else {
			problem.Title = "未知标题"
		}
	}

	// 4. 提取题目难度（映射中英文）
	difficultyMap := map[string]string{
		"Easy":   "Easy",
		"Medium": "Medium",
		"Hard":   "Hard",
		"简单":     "Easy",
		"中等":     "Medium",
		"困难":     "Hard",
	}
	difficultyRegex := regexp.MustCompile(`(Easy|Medium|Hard|简单|中等|困难)`)
	difficultyMatches := difficultyRegex.FindStringSubmatch(fileContent)
	if difficultyMatches != nil {
		problem.Difficulty = difficultyMap[difficultyMatches[1]]
	} else {
		problem.Difficulty = "未知"
	}

	// 5. 提取题目URL
	urlRegex := regexp.MustCompile(`https://leetcode\.cn/problems/[^(\n)]+`)
	urlMatches := urlRegex.FindStringSubmatch(fileContent)
	if urlMatches != nil {
		problem.URL = strings.TrimSpace(urlMatches[0])
	} else {
		// 兜底生成空URL
		problem.URL = ""
	}

	// 6. 提取解法描述
	// 优先匹配 // 解法：xxx 或 // 思路：xxx
	solutionRegex1 := regexp.MustCompile(`//\s*(解法|思路)[:：]\s*([^\n]+)`)
	solutionMatches1 := solutionRegex1.FindStringSubmatch(fileContent)
	if solutionMatches1 != nil && len(solutionMatches1) >= 3 {
		problem.Solution = strings.TrimSpace(solutionMatches1[2])
	} else {
		// 备用：提取核心算法关键词
		algorithmRegex := regexp.MustCompile(`//\s*(栈|递归|动态规划|贪心|二叉堆|前缀和|字典树|LRU|DP|BFS|DFS)[^(\n)]+`)
		algorithmMatches := algorithmRegex.FindStringSubmatch(fileContent)
		if algorithmMatches != nil {
			problem.Solution = algorithmMatches[1]
		} else {
			problem.Solution = ""
		}
	}

	return problem, nil
}

// 生成Markdown表格
func generateMarkdownTable(problems []LeetCodeProblem) string {
	var table strings.Builder
	// 表格头部
	table.WriteString("| 序号 | 题目 | 难度 | 解法 |\n")
	table.WriteString("|------|------|------|------|\n")

	// 表格行
	for _, p := range problems {
		titleCell := fmt.Sprintf("[%s](%s)", p.Title, p.URL)
		table.WriteString(fmt.Sprintf("| %d | %s | %s | %s |\n", p.Num, titleCell, p.Difficulty, p.Solution))
	}

	return table.String()
}

// 更新README中的表格
func updateReadme(tableContent string) error {
	// 读取README内容
	readmeContent, err := os.ReadFile(readmePath)
	if err != nil {
		return fmt.Errorf("读取README失败: %w", err)
	}

	// 替换标记间的内容
	contentStr := string(readmeContent)
	startIdx := strings.Index(contentStr, tableStartMarker)
	endIdx := strings.Index(contentStr, tableEndMarker)
	if startIdx == -1 || endIdx == -1 || startIdx > endIdx {
		return fmt.Errorf("未找到表格标记（%s / %s）", tableStartMarker, tableEndMarker)
	}

	// 拼接新内容
	newContent := fmt.Sprintf(
		"%s\n%s\n%s",
		contentStr[:startIdx+len(tableStartMarker)],
		tableContent,
		contentStr[endIdx:],
	)

	// 写入README
	err = os.WriteFile(readmePath, []byte(newContent), 0644)
	if err != nil {
		return fmt.Errorf("写入README失败: %w", err)
	}

	return nil
}

func main() {
	// 1. 解析所有题目
	problems, err := parseAllProblems()
	if err != nil {
		fmt.Printf("解析题目失败: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("成功解析 %d 道题目\n", len(problems))

	// 2. 生成表格
	table := generateMarkdownTable(problems)

	// 3. 更新README
	err = updateReadme(table)
	if err != nil {
		fmt.Printf("更新README失败: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("✅ README表格更新成功！")
}
