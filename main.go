package main

import (
	"os"

	"github.com/RaymondCode/simple-demo/repository"
	"github.com/gin-gonic/gin"
)

func main() {
	//初始化数据库
	if err := repository.Init(); err != nil {
		os.Exit(-1)
	}

	//创建路由
	r := gin.Default()
	initRouter(r)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
