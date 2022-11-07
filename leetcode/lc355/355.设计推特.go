/*
 * @lc app=leetcode.cn id=355 lang=golang
 *
 * [355] 设计推特
 *
 * https://leetcode.cn/problems/design-twitter/description/
 *
 * algorithms
 * Medium (40.96%)
 * Likes:    320
 * Dislikes: 0
 * Total Accepted:    38.9K
 * Total Submissions: 95K
 * Testcase Example:  '["Twitter","postTweet","getNewsFeed","follow","postTweet","getNewsFeed","unfollow","getNewsFeed"]\n' +
  '[[],[1,5],[1],[1,2],[2,6],[1],[1,2],[1]]'
 *
 * 设计一个简化版的推特(Twitter)，可以让用户实现发送推文，关注/取消关注其他用户，能够看见关注人（包括自己）的最近 10 条推文。
 *
 * 实现 Twitter 类：
 *
 *
 * Twitter() 初始化简易版推特对象
 * void postTweet(int userId, int tweetId) 根据给定的 tweetId 和 userId
 * 创建一条新推文。每次调用此函数都会使用一个不同的 tweetId 。
 * List<Integer> getNewsFeed(int userId) 检索当前用户新闻推送中最近  10 条推文的 ID
 * 。新闻推送中的每一项都必须是由用户关注的人或者是用户自己发布的推文。推文必须 按照时间顺序由最近到最远排序 。
 * void follow(int followerId, int followeeId) ID 为 followerId 的用户开始关注 ID 为
 * followeeId 的用户。
 * void unfollow(int followerId, int followeeId) ID 为 followerId 的用户不再关注 ID 为
 * followeeId 的用户。
 *
 *
 *
 *
 * 示例：
 *
 *
 * 输入
 * ["Twitter", "postTweet", "getNewsFeed", "follow", "postTweet",
 * "getNewsFeed", "unfollow", "getNewsFeed"]
 * [[], [1, 5], [1], [1, 2], [2, 6], [1], [1, 2], [1]]
 * 输出
 * [null, null, [5], null, null, [6, 5], null, [5]]
 *
 * 解释
 * Twitter twitter = new Twitter();
 * twitter.postTweet(1, 5); // 用户 1 发送了一条新推文 (用户 id = 1, 推文 id = 5)
 * twitter.getNewsFeed(1);  // 用户 1 的获取推文应当返回一个列表，其中包含一个 id 为 5 的推文
 * twitter.follow(1, 2);    // 用户 1 关注了用户 2
 * twitter.postTweet(2, 6); // 用户 2 发送了一个新推文 (推文 id = 6)
 * twitter.getNewsFeed(1);  // 用户 1 的获取推文应当返回一个列表，其中包含两个推文，id 分别为 -> [6, 5] 。推文
 * id 6 应当在推文 id 5 之前，因为它是在 5 之后发送的
 * twitter.unfollow(1, 2);  // 用户 1 取消关注了用户 2
 * twitter.getNewsFeed(1);  // 用户 1 获取推文应当返回一个列表，其中包含一个 id 为 5 的推文。因为用户 1
 * 已经不再关注用户 2
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= userId, followerId, followeeId <= 500
 * 0 <= tweetId <= 10^4
 * 所有推特的 ID 都互不相同
 * postTweet、getNewsFeed、follow 和 unfollow 方法最多调用 3 * 10^4 次
 *
 *
*/
package lc355

import (
    "sort"
    "time"
)

// @lc code=start
type Tweet struct {
    id        int
    timestamp int64
    next      *Tweet
}

func newTweet(id int) *Tweet {
    return &Tweet{
        id:        id,
        timestamp: time.Now().UnixNano(),
    }
}

type User struct {
    id            int
    tweetHead     *Tweet
    followedUsers map[*User]struct{}
}

func newUser(id int) *User {
    return &User{
        id:            id,
        tweetHead:     nil,
        followedUsers: map[*User]struct{}{},
    }
}

func (u *User) follow(other *User) {
    u.followedUsers[other] = struct{}{}
}

func (u *User) unfollow(other *User) {
    delete(u.followedUsers, other)
}

func (u *User) post(tweet *Tweet) {
    // 把新发布的tweet作为头结点
    tweet.next = u.tweetHead
    u.tweetHead = tweet
}

type Twitter struct {
    idToUser map[int]*User
}

func Constructor() Twitter {
    return Twitter{
        idToUser: make(map[int]*User),
    }
}

func (t *Twitter) PostTweet(userId int, tweetId int) {
    if _, ok := t.idToUser[userId]; !ok {
        t.idToUser[userId] = newUser(userId)
    }
    t.idToUser[userId].post(newTweet(tweetId))
}

func (t *Twitter) GetNewsFeed(userId int) []int {
    user := t.idToUser[userId]
    res := []int{}
    if user == nil {
        return res
    }
    // 按时间线返回最近10条tweet，合并多个链表
    tweets := make([]*Tweet, 0)
    for tmp := user.tweetHead; tmp != nil; tmp = tmp.next {
        tweets = append(tweets, tmp)
    }
    for other := range user.followedUsers {
        for tmp := other.tweetHead; tmp != nil; tmp = tmp.next {
            tweets = append(tweets, tmp)
        }
    }

    sort.Slice(tweets, func(i, j int) bool {
        return tweets[i].timestamp > tweets[j].timestamp
    })

    for i := 0; i < 10 && i < len(tweets); i++ {
        res = append(res, tweets[i].id)
    }
    return res
}

func (t *Twitter) Follow(followerId int, followeeId int) {
    if _, ok := t.idToUser[followerId]; !ok {
        t.idToUser[followerId] = newUser(followerId)
    }
    if _, ok := t.idToUser[followeeId]; !ok {
        t.idToUser[followeeId] = newUser(followeeId)
    }

    t.idToUser[followerId].follow(t.idToUser[followeeId])
}

func (t *Twitter) Unfollow(followerId int, followeeId int) {
    if _, ok := t.idToUser[followerId]; !ok {
        return
    }
    if _, ok := t.idToUser[followeeId]; !ok {
        return
    }
    t.idToUser[followerId].unfollow(t.idToUser[followeeId])
}

/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */
// @lc code=end
