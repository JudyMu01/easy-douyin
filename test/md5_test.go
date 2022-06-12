package test

import (
	"fmt"
	"testing"

	"github.com/JudyMu01/easy-douyin/util"
)

func TestMd5V(t *testing.T) {
	str := "123456"

	fmt.Println(util.Md5V(str))
}
