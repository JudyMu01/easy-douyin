package service

import (
	"github.com/JudyMu01/easy-douyin/repository"
	"github.com/JudyMu01/easy-douyin/util"
)

type UserData struct {
	Id            int64  `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	FollowCount   int64  `json:"follow_count,omitempty"`
	FollowerCount int64  `json:"follower_count,omitempty"`
	IsFollow      bool   `json:"is_follow,omitempty"`
	LikedCount    int64  `json:"total_favorited,omitempty"`
	LikeCount     int64  `json:"favorite_count,omitempty"`
}

func QueryUserData(userid int64, token string) (*UserData, error) {
	user, err := repository.NewUserDaoInstance().QueryUserById(userid)
	if err != nil {
		util.Logger.Error("find user by id err:" + err.Error())
		return nil, err
	}
	var isFollow bool
	var res int
	if token == "" {
		res, err = repository.QueryFollowInfo(userid, userid)
	} else {
		res, err = repository.QueryFollowInfo(repository.UsersLoginInfo[token].Id, userid)
	}

	if err != nil {
		util.Logger.Error("query follow info err:" + err.Error())
		return nil, err
	} else {
		if res == 1 {
			isFollow = true
		} else {
			isFollow = false
		}
	}
	followCount, _ := repository.CountFollow(userid)
	followerCount, _ := repository.CountFollower(userid)
	likeCount, _ := repository.CountTotalLikeByUser(userid)
	recvCount, _ := countRecvLike(userid)
	var userData = UserData{Id: user.Id, Name: user.Name, FollowCount: followCount, FollowerCount: followerCount, IsFollow: isFollow, LikedCount: recvCount, LikeCount: likeCount}

	return &userData, nil
}

func UserReigster(username string, password string) (*UserData, error) {
	user, err := repository.NewUserDaoInstance().AddUser(username, password)
	if err != nil {
		util.Logger.Error("add user to db err:" + err.Error())
		return nil, err
	}
	var userData = UserData{Id: user.Id, Name: user.Name, FollowCount: 0, FollowerCount: 0, IsFollow: false}
	return &userData, nil
}

func countRecvLike(userid int64) (int64, error) {
	videos, err := repository.NewVideoDaoInstance().SearchVideoById(userid)
	var recvLikes int64 = 0
	for _, k := range videos {
		likes, _ := repository.CountLike(k.Id)
		recvLikes += likes
	}
	return recvLikes, err
}
