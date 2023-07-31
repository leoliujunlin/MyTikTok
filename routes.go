package main

import (
	"TikTok/controller"
	"TikTok/middleware"
	"github.com/gin-gonic/gin"
)

func CollectRoutes(r *gin.Engine) *gin.Engine {
	// 允许跨域访问
	r.Use(middleware.CORSMiddleware())
	// 用户控制
	userController := controller.NewUserController()
	userRoutes := r.Group("douyin/user/")
	{
		userRoutes.POST("register", userController.Register)                        // 注册
		userRoutes.POST("login", userController.Login)                              // 登录
		userRoutes.Use(middleware.AuthMiddleware()).GET("", userController.GetInfo) //获得用户信息
	}
	return r
}
