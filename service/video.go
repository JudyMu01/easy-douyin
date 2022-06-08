package service

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/JudyMu01/easy-douyin/repository"
	"github.com/JudyMu01/easy-douyin/util"
	"github.com/disintegration/imaging"
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

var serverAddr string = "http://192.168.2.119:8080/"

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

func PostVideo(fileName string, title string, userID int64) (*repository.Video, error) {
	playUrl := serverAddr + "public/videos/" + fileName
	//添加生成视频关键帧并上传到public目录的函数
	_, err := GetSnapshot("./public/videos/"+fileName, "./public/covers/"+fileName, 1)
	if err != nil {
		util.Logger.Error("generate cover err:" + err.Error())
		return nil, err
	}
	coverUrl := serverAddr + "public/covers/" + fileName + ".png"
	newVideo := repository.Video{PlayUrl: playUrl, CoverUrl: coverUrl, Title: title, UserId: userID, CreateTime: time.Now(), CommentCount: 0, FavoriteCount: 0}
	video, err := repository.NewVideoDaoInstance().AddVideo(newVideo)
	if err != nil {
		util.Logger.Error("post video to db err:" + err.Error())
		return nil, err
	}
	return video, nil
}

func GetPublishList(userID int64, token string) ([]VideoData, error) {

	videoList, err := repository.NewVideoDaoInstance().SearchVideoById(userID)
	if err != nil {
		util.Logger.Error("search video by id err:" + err.Error())
		return nil, err
	}
	//prepare VideoData list
	var videoDataList []VideoData
	for _, k := range videoList {
		author, _ := QueryUserData(userID, token)
		//isFavorite := QueryLike(k.Id,k.UserId)
		videoData := VideoData{Id: k.Id, Author: *author, PlayUrl: k.PlayUrl, CoverUrl: k.CoverUrl, FavoriteCount: k.FavoriteCount, CommentCount: k.CommentCount, IsFavorite: false, Title: k.Title}
		videoDataList = append(videoDataList, videoData)
	}

	return videoDataList, nil
}

func GetSnapshot(videoPath, snapshotPath string, frameNum int) (snapshotName string, err error) {
	buf := bytes.NewBuffer(nil)
	err = ffmpeg.Input(videoPath).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()
	if err != nil {
		util.Logger.Error("generate cover fail:" + err.Error())
		return "", err
	}

	img, err := imaging.Decode(buf)
	if err != nil {
		util.Logger.Error("decoding cover fail:" + err.Error())
		return "", err
	}

	err = imaging.Save(img, snapshotPath+".png")
	if err != nil {
		util.Logger.Error("saving cover fail:" + err.Error())
		return "", err
	}
	names := strings.Split(snapshotPath, "\\")
	snapshotName = names[len(names)-1] + ".png"
	return snapshotName, nil
}
