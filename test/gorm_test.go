package test

import (
	"fmt"
	"testing"

	"github.com/JudyMu01/easy-douyin/repository"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestGorm(t *testing.T) {
	dsn := "root:piper_2021%wii@tcp(127.0.0.1:3306)/douyin?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		t.Fatal(err)
	}

	data := make([]*repository.User, 0)
	err = db.Find(&data).Error

	if err != nil {
		t.Fatal(err)
	}

	for _, v := range data {
		fmt.Printf("User==> %v\n", v)
	}
}
