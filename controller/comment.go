package controller

import (
	"net/http"
	"strconv"

	"github.com/JudyMu01/easy-douyin/repository"
	"github.com/JudyMu01/easy-douyin/service"
	"github.com/gin-gonic/gin"
)

type CommentListResponse struct {
	Response
	CommentList []service.CommentData `json:"comment_list,omitempty"`
}

type CommentActionResponse struct {
	Response
	Comment service.CommentData `json:"comment"`
}

// CommentAction no practical effect, just check if token is valid
func CommentAction(c *gin.Context) {
	token := c.Query("token")
	actionType := c.Query("action_type")
	videoId, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)

	if user, exist := repository.UsersLoginInfo[token]; exist {
		if actionType == "1" {
			//post a comment
			content := c.Query("comment_text")
			commentData, err := service.PostComment(user.Id, videoId, content)
			if err != nil {
				c.JSON(http.StatusOK, Response{StatusCode: 2, StatusMsg: "post comment fail"})
			} else {
				c.JSON(http.StatusOK, CommentActionResponse{
					Response: Response{StatusCode: 0},
					Comment:  *commentData,
				})
			}
		} else {
			//delete a comment
			commentId, _ := strconv.ParseInt(c.Query("comment_id"), 10, 64)
			commentData, err := service.DelComment(commentId)
			if err == nil {
				c.JSON(http.StatusOK, CommentActionResponse{
					Response: Response{StatusCode: 0},
					Comment:  *commentData,
				})
			} else {
				c.JSON(http.StatusOK, Response{StatusCode: 3, StatusMsg: "delete comment wrong"})
			}
		}
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// CommentList all videos have same demo comment list
func CommentList(c *gin.Context) {
	token := c.Query("token")
	videoId, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
	if _, exist := repository.UsersLoginInfo[token]; exist {
		commentDataList, err := service.GetCommentList(videoId)
		if err != nil {
			c.JSON(http.StatusOK, CommentListResponse{
				Response: Response{StatusCode: 2, StatusMsg: "get comment list err"},
			})
		} else {
			c.JSON(http.StatusOK, CommentListResponse{
				Response:    Response{StatusCode: 0},
				CommentList: commentDataList,
			})
		}

	} else {
		c.JSON(http.StatusOK, CommentListResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}

}
