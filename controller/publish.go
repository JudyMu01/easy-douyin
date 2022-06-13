package controller

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/JudyMu01/easy-douyin/middleware"
	"github.com/JudyMu01/easy-douyin/repository"
	"github.com/JudyMu01/easy-douyin/service"
	"github.com/gin-gonic/gin"
)

type VideoListResponse struct {
	Response
	VideoList []service.VideoData `json:"video_list"`
}

// Publish check token then save upload file to public directory, and save data in db
func Publish(c *gin.Context) {
	token := c.PostForm("token")
	title := c.PostForm("title")
	auth, _ := middleware.ParseToken(token)
	token = auth.GetToken()

	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	filename := filepath.Base(data.Filename)
	user := repository.UsersLoginInfo[token]
	fmt.Println(user.Password)
	finalName := fmt.Sprintf("%d_%s", user.Id, filename)
	saveFile := filepath.Join("./public/videos/", finalName)
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	//传一个视频名字，一个title，一个userID
	newVid, err := service.PostVideo(finalName, title, user.Id)
	if newVid == nil || err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 2,
			StatusMsg:  "save post video to db fail",
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  finalName + " uploaded successfully",
	})
}

// PublishList , get list of videos posted by current login user
func PublishList(c *gin.Context) {
	token := c.Query("token")
	userID, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
	auth, _ := middleware.ParseToken(token)
	token = auth.GetToken()
	videoList, err := service.GetPublishList(userID, token)
	if err != nil {
		fmt.Printf("get publish list failed: %s", err)
		c.JSON(http.StatusOK, VideoListResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  "get publish list failed",
			},
		})
	} else {
		c.JSON(http.StatusOK, VideoListResponse{
			Response: Response{
				StatusCode: 0,
			},
			VideoList: videoList,
		})
	}

}
