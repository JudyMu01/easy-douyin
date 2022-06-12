package repository

import (
	"sync"
	"time"
)

type Comment struct {
	Id         int64     `gorm:"column:comment_id"`
	UserId     int64     `gorm:"column:user_id"`
	VideoId    int64     `gorm:"column:video_id"`
	Content    string    `gorm:"column:content"`
	CreateTime time.Time `gorm:"column:create_time"`
}

func (Comment) TableName() string {
	return "comment"
}

type CommentDao struct {
}

var commentDao *CommentDao
var commentOnce sync.Once

func NewCommentDaoInstance() *CommentDao {
	commentOnce.Do(
		func() {
			commentDao = &CommentDao{}
		})
	return commentDao
}

// add a comment
func (*CommentDao) CreateComment(comment *Comment) (*Comment, error) {
	var preComment Comment
	db.Last(&preComment)
	comment.Id = preComment.Id + 1
	err := db.Create(comment).Error

	return comment, err
}

// delete a comment
func (*CommentDao) DeleteComment(commentId int64) (*Comment, error) {

	var comment Comment
	err := db.Model(&Comment{}).Where("comment_id = ?", commentId).Delete(&comment).Error

	return &comment, err
}

// get comment list of a video
func (*CommentDao) CommentList(videoId int64) ([]Comment, error) {
	var comments []Comment
	err := db.Model(&Comment{}).Where("video_id = ?", videoId).Find(&comments).Error

	return comments, err
}

// count comment number
func CountComment(videoId int64) (int64, error) {
	var count int64
	err := db.Model(&Comment{}).Where("video_id = ?", videoId).Count(&count).Error

	return count, err
}
