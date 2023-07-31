package controller

import (
	"TikTok/common"
	"TikTok/model"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
)

type UserController struct {
	DB *gorm.DB
}

func NewUserController() UserController {
	return UserController{DB: common.GetDB()}
}

func (a UserController) Register(c *gin.Context) {
	// 获取参数
	userName, ok := c.GetPostForm("username")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"status_code": 1,
			"status_msg":  "参数名有误",
		})
		return
	}
	password, ok := c.GetPostForm("password")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"status_code": 1,
			"status_msg":  "参数名有误",
		})
		return
	}

	// 数据验证
	var user model.User
	a.DB.Where("name = ?", userName).First(&user)
	if user.ID != 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status_code": 1,
			"status_msg":  "该用户名已存在",
		})
		return
	}
	// 密码加密
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	// 创建用户
	newUser := model.User{
		Name:     userName,
		Password: string(hashedPassword),
	}
	// 发放token
	token, err := common.ReleaseToken(newUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status_code": 1,
			"status_msg":  "系统错误",
		})
		return
	}
	a.DB.Create(&newUser)
	// 返回结果
	c.JSON(http.StatusOK, gin.H{
		"status_code": 0, // 接口约定,0表示成功
		"status_msg":  "注册成功",
		"user_id":     newUser.ID,
		"token":       token,
	})
}

func (a UserController) Login(c *gin.Context) {
	// 获取参数
	userName, ok := c.GetPostForm("username")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"status_code": 1,
			"status_msg":  "参数名有误",
		})
		return
	}
	password, ok := c.GetPostForm("password")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"status_code": 1,
			"status_msg":  "参数名有误",
		})
		return
	}

	// 数据验证
	var user model.User
	// 验证是否存在该用户
	a.DB.Where("name = ?", userName).First(&user)
	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status_code": 1,
			"status_msg":  "该用户名不存在",
		})
		return
	}
	// 判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status_code": 1,
			"status_msg":  "密码错误",
		})
		return
	}
	// 发放token
	token, err := common.ReleaseToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status_code": 1,
			"status_msg":  "系统错误",
		})
		return
	}
	// 成功发送
	c.JSON(http.StatusOK, gin.H{
		"status_code": 0, // 接口约定,0表示成功
		"status_msg":  "登录成功",
		"user_id":     user.ID,
		"token":       token,
	})
}

func (a UserController) GetInfo(c *gin.Context) {
	// 获取上下文中的用户信息
	k, _ := c.Get("user")
	user := k.(model.User)
	c.JSON(http.StatusOK, gin.H{
		"status_code": 0,
		"status_msg":  "获取信息成功",
		"user": gin.H{
			"id":             user.ID,
			"name":           user.Name,
			"follow_count":   user.FollowerCount,
			"follower_count": user.FollowerCount,
			// TODO:这里暂时写为true,需要再改
			"is_follow": true,
		},
	})
}
