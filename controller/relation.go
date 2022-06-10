package controller

import (
	"net/http"
	"strconv"

	"github.com/JudyMu01/easy-douyin/repository"
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

	if _, exist := repository.UsersLoginInfo[token]; exist {
		var err error
		if actionType == "1" {
			//follow
			err = service.AddFollow(token, toUserId)

		} else {
			//unfollow
			err = service.CancelFollow(token, toUserId)
		}
		if err == nil {
			c.JSON(http.StatusOK, Response{StatusCode: 0})
		} else {
			c.JSON(http.StatusOK, Response{StatusCode: 2, StatusMsg: "update follow db fail"})
		}

	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// FollowList all users have same follow list
func FollowList(c *gin.Context) {
	token := c.Query("token")
	userId, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
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
