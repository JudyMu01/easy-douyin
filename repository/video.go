package repository

import (
	"fmt"
	"sync"
	"time"

	"github.com/JudyMu01/easy-douyin/util"
	"gorm.io/gorm"
)

type Video struct {
	Id         int64     `gorm:"column:video_id"`
	PlayUrl    string    `gorm:"column:play_url"`
	CoverUrl   string    `gorm:"column:cover_url"`
	Title      string    `gorm:"column:title"`
	CreateTime time.Time `gorm:"column:create_time"`
	UserId     int64     `gorm:"column:user_id"`
}

func (Video) TableName() string {
	return "video"
}

type VideoDao struct {
}

var videoDao *VideoDao
var videoOnce sync.Once

func NewVideoDaoInstance() *VideoDao {
	videoOnce.Do(
		func() {
			videoDao = &VideoDao{}
		})
	return videoDao
}

// feed videos before latest time, order by create_time
func (*VideoDao) VideoPrepare(latest int64) ([]*Video, error) {

	var videoList []*Video
	//format the time to compare with the time in db
	tm := time.Unix(latest/1000, 0)
	err := db.Model(&Video{}).Where("create_time < ?", tm).Order("create_time desc").Find(&videoList).Error
	//err := db.Model(&Video{}).Find(&videoList).Error
	fmt.Println("length of video list: ", len(videoList))
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		util.Logger.Error("prepare videos err:" + err.Error())
		return nil, err
	}

	return videoList, nil
}

// create a new video in db
func (*VideoDao) AddVideo(newVideo Video) (*Video, error) {
	var oldVideo Video
	db.Last(&oldVideo) //video that has the max id
	newVideo.Id = oldVideo.Id + 1
	err := db.Create(&newVideo).Error
	if err != nil {
		util.Logger.Error("add video err:" + err.Error())
		return nil, err
	}

	return &newVideo, nil
}

// get a publish list of the login user
func (*VideoDao) SearchVideoById(userID int64) ([]*Video, error) {
	var videoList []*Video
	err := db.Model(&Video{}).Where("user_id = ?", userID).Find(&videoList).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		util.Logger.Error("search user videos err:" + err.Error())
		return nil, err
	}
	return videoList, nil
}

// get userId of a video
func (*VideoDao) SearchUserByVid(videoID int64) (*Video, error) {
	var video Video
	err := db.Model(&Video{}).Where("video_id = ?", videoID).Find(&video).Error
	if err != nil {
		util.Logger.Error("search video's author id err:" + err.Error())
		return nil, err
	}
	return &video, nil
}

// count like number
func CountLike(videoId int64) (int64, error) {
	var count int64
	err := db.Model(&Like{}).Where("video_id = ?", videoId).Count(&count).Error
	if err != nil {
		util.Logger.Error("count like err: " + err.Error())
		return -1, err
	}
	return count, nil
}
