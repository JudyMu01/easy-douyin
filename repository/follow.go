package repository

import (
	"sync"
	"time"

	"github.com/JudyMu01/easy-douyin/util"
)

type Follow struct {
	FromUserId int64     `gorm:"column:user_id"`
	ToUserId   int64     `gorm:"column:to_user_id"`
	CreateTime time.Time `gorm:"column:create_time"`
}

func (Follow) TableName() string {
	return "follow"
}

type FollowDao struct {
}

var followDao *FollowDao
var followOnce sync.Once

func NewFollowDaoInstance() *FollowDao {
	followOnce.Do(
		func() {
			followDao = &FollowDao{}
		})
	return followDao
}

// add a follow record
func (*FollowDao) CreateFollow(follow Follow) error {

	err := db.Create(&follow).Error

	return err
}

// delete a follow record
func (*FollowDao) DeleteFollow(fromId int64, toId int64) error {
	// delete follow record in follows
	var follow Follow
	err := db.Where("user_id = ? && to_user_id = ?", fromId, toId).Delete(&follow).Error

	return err
}

// find whether there exist a follow relation between from and to user
func QueryFollowInfo(fromId int64, toId int64) (int, error) {
	var count int64 = 0
	err := db.Model(&Follow{}).Where(Follow{FromUserId: fromId, ToUserId: toId}).Count(&count).Error
	if err != nil {
		util.Logger.Error("query relation err:" + err.Error())
		return 0, err
	}
	if count == 0 {
		return 2, nil
	} else {
		return 1, nil
	}
}

// count user follow number
func CountFollow(userId int64) (int64, error) {
	var count int64
	err := db.Model(&Follow{}).Where("user_id = ?", userId).Count(&count).Error
	if err != nil {
		util.Logger.Error("query follow count err:" + err.Error())
		return 0, err
	}
	return count, nil
}

// count user follower number
func CountFollower(userId int64) (int64, error) {
	var count int64
	err := db.Model(&Follow{}).Where("to_user_id = ?", userId).Count(&count).Error
	if err != nil {
		util.Logger.Error("query follower count err:" + err.Error())
		return 0, err
	}
	return count, nil
}

// get follow list
func GetFollowList(userId int64) ([]int64, error) {
	var followList []int64
	err := db.Model(&Follow{}).Where("user_id = ?", userId).Select("to_user_id").Find(&followList).Error
	if err != nil {
		util.Logger.Error("query follow list err:" + err.Error())
		return nil, err
	}
	return followList, nil
}

// get follower list
func GetFollowerList(toUserId int64) ([]int64, error) {
	var followerList []int64
	err := db.Model(&Follow{}).Where("to_user_id = ?", toUserId).Select("user_id").Find(&followerList).Error
	if err != nil {
		util.Logger.Error("query follower list err:" + err.Error())
		return nil, err
	}
	return followerList, nil
}
