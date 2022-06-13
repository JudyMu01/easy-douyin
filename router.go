package main

import (
	"net/http"

	"github.com/JudyMu01/easy-douyin/controller"
	"github.com/JudyMu01/easy-douyin/middleware"
	"github.com/gin-gonic/gin"
)

func initRouter(r *gin.Engine) {

	auth := middleware.JWY()
	// public directory is used to serve static resources
	r.StaticFS("/public", http.Dir("./public"))
	apiRouter := r.Group("/douyin")

	// basic apis
	apiRouter.GET("/feed/", controller.Feed)
	apiRouter.GET("/user/", auth, controller.UserInfo)
	apiRouter.POST("/user/register/", controller.Register)
	apiRouter.POST("/user/login/", controller.Login)
	apiRouter.POST("/publish/action/", auth, controller.Publish)
	apiRouter.GET("/publish/list/", auth, controller.PublishList)

	// extra apis - I
	apiRouter.POST("/favorite/action/", auth, controller.FavoriteAction)
	apiRouter.GET("/favorite/list/", auth, controller.FavoriteList)
	apiRouter.POST("/comment/action/", auth, controller.CommentAction)
	apiRouter.GET("/comment/list/", controller.CommentList)

	// extra apis - II
	apiRouter.POST("/relation/action/", auth, controller.RelationAction)
	apiRouter.GET("/relation/follow/list/", auth, controller.FollowList)
	apiRouter.GET("/relation/follower/list/", auth, controller.FollowerList)
}
