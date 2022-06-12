package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/JudyMu01/easy-douyin/service"
	"github.com/gin-gonic/gin"
)

//response类型
type FeedResponse struct {
	Response
	VideoList []service.VideoData `json:"video_list,omitempty"`
	NextTime  int64               `json:"next_time,omitempty"`
}

// Feed a list for 10 videos for every request
func Feed(c *gin.Context) {
	latest_time, _ := strconv.ParseInt(c.Query("latest_time"), 10, 64)
	fmt.Println(latest_time)
	token := c.Query("token")
	videos, nextTime, err := service.PrepareVideoData(latest_time, token)
	if err == nil {
		c.JSON(http.StatusOK, FeedResponse{
			Response:  Response{StatusCode: 0},
			VideoList: videos,
			NextTime:  nextTime,
		})
	} else {
		c.JSON(http.StatusOK, FeedResponse{
			Response: Response{StatusCode: 0, StatusMsg: "prepare video failure."},
		})
	}

}
