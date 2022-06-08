package service

import (
	"github.com/JudyMu01/easy-douyin/repository"
	"github.com/JudyMu01/easy-douyin/util"
)

var UsersLoginInfo = map[string]UserData{
	"zhangleidouyin": {
		Id:            1,
		Name:          "zhanglei",
		FollowCount:   10,
		FollowerCount: 5,
		IsFollow:      true,
	},
}

type UserData struct {
	Id            int64  `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	FollowCount   int64  `json:"follow_count,omitempty"`
	FollowerCount int64  `json:"follower_count,omitempty"`
	IsFollow      bool   `json:"is_follow,omitempty"`
}

func QueryUserData(userid int64, token string) (*UserData, error) {
	user, err := repository.NewUserDaoInstance().QueryUserById(userid)
	if err != nil {
		util.Logger.Error("find user by id err:" + err.Error())
		return nil, err
	}
	var userData = UserData{Id: user.Id, Name: user.Name, FollowCount: user.Follow_count, FollowerCount: user.Follower_count, IsFollow: true}

	return &userData, nil
}

func UserReigster(username string, password string) (*UserData, error) {
	user, err := repository.NewUserDaoInstance().AddUser(username, password)
	if err != nil {
		util.Logger.Error("add user to db err:" + err.Error())
		return nil, err
	}
	var userData = UserData{Id: user.Id, Name: user.Name, FollowCount: user.Follow_count, FollowerCount: user.Follower_count, IsFollow: false}
	return &userData, nil
}
