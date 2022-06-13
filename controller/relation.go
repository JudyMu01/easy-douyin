package controller

import (
	"net/http"
	"strconv"

	"github.com/JudyMu01/easy-douyin/middleware"
	"github.com/JudyMu01/easy-douyin/service"
	"github.com/gin-gonic/gin"
)

type UserListResponse struct {
	Response
	UserList []service.UserData `json:"user_list"`
}

// RelationAction no practical effect, just check if token is valid
func RelationAction(c *gin.Context) {
	token := c.Query("token")
	toUserId, _ := strconv.ParseInt(c.Query("to_user_id"), 10, 64)
	actionType := c.Query("action_type") // 1 or 2
	auth, _ := middleware.ParseToken(token)
	token = auth.GetToken()
	var err error
	if actionType == "1" {
		err = service.AddFollow(token, toUserId)
	} else {
		err = service.CancelFollow(token, toUserId)
	}
	if err == nil {
		c.JSON(http.StatusOK, Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "update follow db fail"})
	}

}

// FollowList all users have same follow list
func FollowList(c *gin.Context) {
	token := c.Query("token")
	userId, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
	auth, _ := middleware.ParseToken(token)
	token = auth.GetToken()
	userDataList, err := service.FollowList(userId, token)
	if err != nil {
		c.JSON(http.StatusOK, UserListResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  "service get follow list err",
			},
		})
	}
	c.JSON(http.StatusOK, UserListResponse{
		Response: Response{
			StatusCode: 0,
		},
		UserList: userDataList,
	})
}

// FollowerList all users have same follower list
func FollowerList(c *gin.Context) {
	token := c.Query("token")
	userId, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
	auth, _ := middleware.ParseToken(token)
	token = auth.GetToken()
	userDataList, err := service.FollowerList(userId, token)
	if err != nil {
		c.JSON(http.StatusOK, UserListResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  "service get follow list err",
			},
		})
	}
	c.JSON(http.StatusOK, UserListResponse{
		Response: Response{
			StatusCode: 0,
		},
		UserList: userDataList,
	})
}
