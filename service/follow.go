package service

import (
	"time"

	"github.com/JudyMu01/easy-douyin/repository"
	"github.com/JudyMu01/easy-douyin/util"
)

func AddFollow(token string, toUserId int64) error {

	var follow = repository.Follow{FromUserId: repository.UsersLoginInfo[token].Id, ToUserId: toUserId, CreateTime: time.Now()}
	newFollow, err := repository.NewFollowDaoInstance().CreateFollow(follow)
	if newFollow == nil {
		util.Logger.Error("add follow fail")
		return nil
	}
	if err != nil {
		util.Logger.Error("add follow record err:" + err.Error())
		return err
	}
	return nil
}
func CancelFollow(token string, toUserId int64) error {

	userId := repository.UsersLoginInfo[token].Id
	err := repository.NewFollowDaoInstance().DeleteFollow(userId, toUserId)
	if err != nil {
		util.Logger.Error("delete follow record err:" + err.Error())
		return err
	}

	return nil
}

// return a list of UserData that login user follows
func FollowList(userId int64, token string) ([]UserData, error) {
	userIds, err := repository.GetFollowList(userId)
	if err != nil {
		util.Logger.Error("query follow list err:" + err.Error())
		return nil, err
	}
	userDataList, _ := prepareUserData(userIds, token)
	return userDataList, nil
}

func FollowerList(userId int64, token string) ([]UserData, error) {
	userIds, err := repository.GetFollowerList(userId)
	if err != nil {
		util.Logger.Error("query follower list err:" + err.Error())
		return nil, err
	}
	userDataList, _ := prepareUserData(userIds, token)

	return userDataList, nil
}

// used only in prepare UserData for FollowList and FollowerList
func prepareUserData(userIds []int64, token string) ([]UserData, error) {
	var userDataList []UserData
	for _, k := range userIds {
		m, err := QueryUserData(k, token)
		if err != nil {
			util.Logger.Error("query userdata err:" + err.Error())
			return nil, err
		}
		userDataList = append(userDataList, *m)
	}
	return userDataList, nil
}
