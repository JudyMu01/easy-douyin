package controller

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/JudyMu01/easy-douyin/service"
	"github.com/gin-gonic/gin"
)

type VideoListResponse struct {
	Response
	VideoList []service.VideoData `json:"video_list"`
}

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	token := c.PostForm("token")

	if _, exist := service.UsersLoginInfo[token]; !exist {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}

	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	filename := filepath.Base(data.Filename)
	user := service.UsersLoginInfo[token]
	finalName := fmt.Sprintf("%d_%s", user.Id, filename)
	saveFile := filepath.Join("./public/", finalName)
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
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
