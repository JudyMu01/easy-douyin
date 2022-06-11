package repository

import (
	"sync"
	"time"

	"github.com/JudyMu01/easy-douyin/util"
)

type Like struct {
	UserId     int64     `gorm:"column:user_id"`
	VideoId    int64     `gorm:"column:video_id"`
	CreateTime time.Time `gorm:"column:create_time"`
}

func (Like) TableName() string {
	return "like"
}

type LikeDao struct {
}

var likeDao *LikeDao
var likeOnce sync.Once

func NewLikeDaoInstance() *LikeDao {
	likeOnce.Do(
		func() {
			likeDao = &LikeDao{}
		})
	return likeDao
}

// add a like record
func (*LikeDao) CreateLike(like Like) error {

	err := db.Create(&like).Error

	return err
}

// delete a like record
func (*LikeDao) DeleteLike(userId int64, videoId int64) error {
	// delete follow record in follows
	var like Like
	err := db.Model(&Like{}).Where("user_id = ? && video_id = ?", userId, videoId).Delete(&like).Error

	return err
}

// query like relation
func QueryLike(videoId int64, userId int64) (bool, error) {
	var count int64 = 0
	err := db.Model(&Like{}).Where(Like{UserId: userId, VideoId: videoId}).Count(&count).Error
	if err != nil {
		util.Logger.Error("query relation err:" + err.Error())
		return false, err
	}
	if count == 0 {
		return false, nil
	} else {
		return true, nil
	}
}

// get a list of videos liked by the login user
func GetFavoriteList(userId int64) ([]int64, error) {
	var videoIds []int64
	err := db.Model(&Like{}).Where("user_id = ?", userId).Select("video_id").Find(&videoIds).Error
	if err != nil {
		util.Logger.Error("get favorite list err: " + err.Error())
		return nil, err
	}
	return videoIds, nil
}
