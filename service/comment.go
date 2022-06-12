package service

import (
	"time"

	"github.com/JudyMu01/easy-douyin/repository"
	"github.com/JudyMu01/easy-douyin/util"
)

type CommentData struct {
	Id         int64    `json:"id,omitempty"`
	User       UserData `json:"user"`
	Content    string   `json:"content,omitempty"`
	CreateDate string   `json:"create_date,omitempty"`
}

func PostComment(userId int64, videoId int64, content string) (*CommentData, error) {

	// save to db
	comment := repository.Comment{UserId: userId, VideoId: videoId, Content: content, CreateTime: time.Now()}
	newComment, err := repository.NewCommentDaoInstance().CreateComment(&comment)
	if err != nil {
		util.Logger.Error("create comment err: " + err.Error())
		return nil, err
	}
	// prepare return structure
	commentData, _ := prepare_CommentData(*newComment)
	return commentData, nil
}

func DelComment(commentId int64) (*CommentData, error) {
	// delete in db
	comment, err := repository.NewCommentDaoInstance().DeleteComment(commentId)
	if err != nil {
		util.Logger.Error("delete comment err: " + err.Error())
		return nil, err
	}
	// prepare return structure
	commentData, _ := prepare_CommentData(*comment)
	return commentData, nil
}

func GetCommentList(videoId int64) ([]CommentData, error) {
	comments, err := repository.NewCommentDaoInstance().CommentList(videoId)
	if err != nil {
		util.Logger.Error("get comment list err:" + err.Error())
		return nil, err
	}
	var commentDataList []CommentData
	for _, k := range comments {
		commentData, _ := prepare_CommentData(k)
		commentDataList = append(commentDataList, *commentData)
	}
	return commentDataList, nil
}

// used by above three functions
func prepare_CommentData(comment repository.Comment) (*CommentData, error) {
	userData, err := QueryUserData(comment.UserId, "")
	if err != nil {
		util.Logger.Error("prepare comment user data err: " + err.Error())
		return nil, err
	}
	commentData := CommentData{Id: comment.Id, User: *userData, Content: comment.Content, CreateDate: comment.CreateTime.Format("01-02")}
	return &commentData, nil
}
