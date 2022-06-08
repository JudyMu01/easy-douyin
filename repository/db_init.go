package repository

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() error {
	var err error
	dsn := "root:piper_2021%wii@tcp(127.0.0.1:3306)/douyin?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	result := make([]*User, 0)
	db.Find(&result)
	for _, i := range result {
		fmt.Printf("User ==> %v \n", i)
	}
	return err

}
