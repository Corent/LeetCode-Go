package algorithm355

import "sort"

// https://blog.csdn.net/m0_54190747/article/details/134498174

type Twitter struct {
	users       []int
	followed    map[int][]int
	tweetposted map[int][]int
	postedTime  map[int]int
	time        int
}

func Constructor() Twitter {
	return Twitter{
		users:       []int{},
		followed:    map[int][]int{},
		tweetposted: map[int][]int{},
		postedTime:  map[int]int{},
		time:        0,
	}
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	this.tweetposted[userId] = append(this.tweetposted[userId], tweetId)
	this.postedTime[tweetId] = this.time
	this.time++
}

func (this *Twitter) GetNewsFeed(userId int) []int {
	post := []int{}
	post = append(post, this.tweetposted[userId]...)
	for _, v := range this.followed[userId] {
		post = append(post, this.tweetposted[v]...)
	}
	sort.Slice(post, func(i, j int) bool {
		return this.postedTime[post[i]] > this.postedTime[post[j]]
	})
	if len(post) > 10 {
		return post[:10]
	}
	return post
}

func (this *Twitter) Follow(followerId int, followeeId int) {
	isfollowed := false
	for _, v := range this.followed[followerId] {
		if v == followeeId {
			isfollowed = true
			break
		}
	}
	if !isfollowed {
		this.followed[followerId] = append(this.followed[followerId], followeeId)
	}
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	x := this.followed[followerId]
	delete(this.followed, followerId)
	for _, v := range x {
		if v != followeeId {
			this.followed[followerId] = append(this.followed[followerId], v)
		}
	}
}

/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */
