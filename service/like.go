package service

import (
	"time"

	"github.com/JudyMu01/easy-douyin/repository"
	"github.com/JudyMu01/easy-douyin/util"
)

func AddLike(token string, videoId int64) error {

	var like = repository.Like{UserId: repository.UsersLoginInfo[token].Id, VideoId: videoId, CreateTime: time.Now()}
	err := repository.NewLikeDaoInstance().CreateLike(like)
	if err != nil {
		util.Logger.Error("add like record err:" + err.Error())
		return err
	}

	return nil
}
func CancelLike(token string, videoId int64) error {

	userId := repository.UsersLoginInfo[token].Id
	err := repository.NewLikeDaoInstance().DeleteLike(userId, videoId)
	if err != nil {
		util.Logger.Error("delete like record err:" + err.Error())
		return err
	}

	return nil
}
func GetFavoriteList(userId int64, token string) ([]VideoData, error) {
	videoIds, _ := repository.GetFavoriteList(userId)
	var videoDataList []VideoData
	for _, k := range videoIds {
		video, _ := repository.NewVideoDaoInstance().SearchUserByVid(k)
		author, _ := QueryUserData(video.UserId, token)
		isFavorite, _ := repository.QueryLike(k, video.UserId)
		favoriteCount, _ := repository.CountLike(k)
		commentCount, _ := repository.CountComment(k)
		videoData := VideoData{Id: k, Author: *author, PlayUrl: video.PlayUrl, CoverUrl: video.CoverUrl, FavoriteCount: favoriteCount, CommentCount: commentCount, IsFavorite: isFavorite, Title: video.Title}
		videoDataList = append(videoDataList, videoData)
	}
	return videoDataList, nil
}
