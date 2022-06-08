package service

import (
	"time"

	"github.com/JudyMu01/easy-douyin/repository"
	"github.com/JudyMu01/easy-douyin/util"
)

type VideoData struct {
	Id            int64    `json:"id,omitempty"`
	Author        UserData `json:"author"`
	PlayUrl       string   `json:"play_url,omitempty"`
	CoverUrl      string   `json:"cover_url,omitempty"`
	FavoriteCount int64    `json:"favorite_count,omitempty"`
	CommentCount  int64    `json:"comment_count,omitempty"`
	IsFavorite    bool     `json:"is_favorite,omitempty"`
	Title         string   `json:"title,omitempty"`
}

//prepare video data list in the response format
func PrepareVideoData(latestTime int64, token string) ([]VideoData, int64, error) {
	//if doesn't get a last_time timestamp, use current time
	if latestTime == 0 {
		latestTime = time.Now().Unix()
	}
	videos, err := repository.NewVideoDaoInstance().VideoPrepare(latestTime)
	if err != nil {
		util.Logger.Error("find video list from VideoPrepare err:" + err.Error())
		return nil, time.Now().Unix(), err
	}
	var videoData = make([]VideoData, 10)
	var next_time int64
	for i, k := range videos {
		author, _ := QueryUserData(k.UserId, token)
		//isFavorite := QueryLike(k.Id,k.UserId)
		videoData[i] = VideoData{Id: k.Id, Author: *author, PlayUrl: k.PlayUrl, CoverUrl: k.CoverUrl, FavoriteCount: k.FavoriteCount, CommentCount: k.CommentCount, IsFavorite: false, Title: k.Title}
		next_time = k.CreateTime.Unix()
	}

	return videoData, next_time, nil
}

func PostVideo(videoData VideoData) (*repository.Video, error) {
	// var oldVideo Video
	// db.Last(&oldVideo) //video that has the max id
	// newVideo := Video{Id: video.Id + 1, PlayUrl: newVideoData.PlayUrl, CoverUrl: newVideoData.CoverUrl, Title: newVideoData.Title, UserId: newVideoData.Author.Id, CreateTime: time.Now(), CommentCount: 0, FavoriteCount: 0}
	return nil, nil
}
