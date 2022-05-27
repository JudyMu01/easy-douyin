package repository

import (
	"sync"
	"time"

	"github.com/RaymondCode/simple-demo/util"
	"gorm.io/gorm"
)

type User struct {
	Id             int64     `gorm:"column:user_id"`
	Name           string    `gorm:"column:username"`
	Password       string    `gorm:"column:password"`
	Follow_count   int64     `gorm:"column:follow_count"`
	Follower_count int64     `gorm:"column:follower_count"`
	CreateTime     time.Time `gorm:"column:create_time"`
	ModifyTime     time.Time `gorm:"column:modify_time"`
}

func (User) TableName() string {
	return "user"
}

type UserDao struct {
}

var userDao *UserDao
var userOnce sync.Once

func NewUserDaoInstance() *UserDao {
	userOnce.Do(
		func() {
			userDao = &UserDao{}
		})
	return userDao
}

func (*UserDao) QueryUserById(id int64) (*User, error) {
	var user User
	err := db.Where("user_id = ?", id).Find(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		util.Logger.Error("find user by id err:" + err.Error())
		return nil, err
	}
	return &user, nil
}

// func (*UserDao) MQueryUserById(ids []int64) (map[int64]*User, error) {
// 	var users []*User
// 	err := db.Where("id in (?)", ids).Find(&users).Error
// 	if err != nil {
// 		util.Logger.Error("batch find user by id err:" + err.Error())
// 		return nil, err
// 	}
// 	userMap := make(map[int64]*User)
// 	for _, user := range users {
// 		userMap[user.Id] = user
// 	}
// 	return userMap, nil
// }

func (*UserDao) QueryUserByName(username string) (*User, error) {
	var user User
	err := db.Where("username = ?", username).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		util.Logger.Error("find user by id err:" + err.Error())
		return nil, err
	}
	return &user, nil
}

func (*UserDao) AddUser(username string, password string) (*User, error) {
	var user User
	db.Last(&user) //max id
	newUser := User{Id: user.Id + 1, Name: username, Password: password, Follow_count: 0, Follower_count: 0, CreateTime: time.Now(), ModifyTime: time.Now()}
	err := db.Create(&newUser).Error
	if err != nil {
		util.Logger.Error("create user err:" + err.Error())
		return nil, err
	}
	return userDao.QueryUserByName(username)
}
