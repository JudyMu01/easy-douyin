package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/RaymondCode/simple-demo/repository"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
)

// usersLoginInfo use map to store user info, and key is username+password for demo

//登陆或者注册结束后返回给客户端
type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

//请求用户信息的返回
type UserResponse struct {
	Response
	User service.UserData `json:"user"`
}

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	token := username + password
	//query if user exists
	var usersLoginInfo, err = repository.NewUserDaoInstance().QueryUserByName(username)
	if usersLoginInfo != nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User already exist"},
		})
	} else if err != nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 2, StatusMsg: "when register,checking existence in db has error"},
		})
	} else {
		//user doesn't exist
		user, err := service.UserReigster(username, password)
		if err != nil {
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: Response{StatusCode: 3, StatusMsg: "add new user error"},
			})
		} else {

		}
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0},
			UserId:   user.Id,
			Token:    token,
		})
	}
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	token := username + password
	var usersLoginInfo, err = repository.NewUserDaoInstance().QueryUserByName(username)

	if err != nil {
		fmt.Println(err)
	}
	if usersLoginInfo.Password != password {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 2, StatusMsg: "Password wrong"},
		})
	} else if usersLoginInfo != nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0},
			UserId:   usersLoginInfo.Id,
			Token:    token,
		})
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}

func UserInfo(c *gin.Context) {
	userid, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
	token := c.Query("token")
	userData, err := service.QueryUserData(userid, token)
	if err == nil {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 0},
			User:     *userData,
		})
	} else {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}
