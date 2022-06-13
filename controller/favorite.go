package controller

import (
	"net/http"
	"strconv"

	"github.com/JudyMu01/easy-douyin/middleware"
	"github.com/JudyMu01/easy-douyin/service"
	"github.com/gin-gonic/gin"
)

// FavoriteAction, like a video or cancel like
func FavoriteAction(c *gin.Context) {
	token := c.Query("token")
	videoId, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
	actionType := c.Query("action_type")
	auth, _ := middleware.ParseToken(token)
	token = auth.GetToken()
	var err error
	if actionType == "1" {
		err = service.AddLike(token, videoId)
	} else {
		err = service.CancelLike(token, videoId)
	}
	if err != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "add or cancel like err"})
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 0})
	}

}

// FavoriteList all users have same favorite video list
func FavoriteList(c *gin.Context) {
	userId, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
	token := c.Query("token")
	claim, _ := middleware.ParseToken(token)
	token = claim.GetToken()
	videoData, err := service.GetFavoriteList(userId, token)
	if err != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "get favorite list err"})
	} else {
		c.JSON(http.StatusOK, VideoListResponse{
			Response: Response{
				StatusCode: 0,
			},
			VideoList: videoData,
		})
	}
}
