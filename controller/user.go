package controller

import (
	"net/http"
	"strconv"

	"github.com/JudyMu01/easy-douyin/repository"
	"github.com/JudyMu01/easy-douyin/service"
	"github.com/JudyMu01/easy-douyin/util"
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
	password = util.Md5V(password)
	token := username + password

	//query if user exists
	if _, exist := repository.UsersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User already exist"},
		})
	} else {
		user, err := service.UserReigster(username, password)
		if err != nil {
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: Response{StatusCode: 2, StatusMsg: "add new user error"},
			})
		} else {
			// update token map
			repository.NewUserDaoInstance().TokenMap()
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: Response{StatusCode: 0},
				UserId:   user.Id,
				Token:    token,
			})
		}
	}
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	password = util.Md5V(password)
	token := username + password

	if user, exist := repository.UsersLoginInfo[token]; exist {
		if repository.UsersLoginInfo[token].Password != password {
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: Response{StatusCode: 2, StatusMsg: "Password wrong"},
			})
		} else {
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: Response{StatusCode: 0},
				UserId:   user.Id,
				Token:    token,
			})
		}
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
