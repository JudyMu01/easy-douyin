package repository

import (
	"sync"
	"time"

	"github.com/JudyMu01/easy-douyin/util"
	"gorm.io/gorm"
)

var UsersLoginInfo map[string]User

type User struct {
	Id         int64     `gorm:"column:user_id"`
	Name       string    `gorm:"column:username"`
	Password   string    `gorm:"column:password"`
	CreateTime time.Time `gorm:"column:create_time"`
	ModifyTime time.Time `gorm:"column:modify_time"`
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

//when initing db, create the token-User map for checking token in following user operations.
func (*UserDao) TokenMap() {
	UsersLoginInfo = make(map[string]User)
	result := make([]*User, 0)
	db.Find(&result)
	for _, i := range result {
		UsersLoginInfo[i.Name+i.Password] = *i
	}
}

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
	newUser := User{Id: user.Id + 1, Name: username, Password: password, CreateTime: time.Now(), ModifyTime: time.Now()}
	err := db.Create(&newUser).Error
	if err != nil {
		util.Logger.Error("create user err:" + err.Error())
		return nil, err
	}
	return userDao.QueryUserByName(username)
}
