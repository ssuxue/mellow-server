package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"memo/model"
	"memo/util"
	"net/http"
	"strconv"
	"time"
)

// Gin 获取 POST JSON DATA 封装登录参数
type LoginForm struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func GetUserById(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))
	userModel := model.User{}
	user := userModel.FindById(id)

	result := model.Result{
		Code:    http.StatusOK,
		Message: "查询成功",
		Data:    user,
	}

	context.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}

func GetUserByPhone(context *gin.Context) {
	phone := context.Param("phone")
	userModel := model.User{}
	user := userModel.FindByPhone(phone)

	result := model.Result{
		Code:    http.StatusOK,
		Message: "查询成功",
		Data:    user,
	}

	context.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}

// 通过用户名注册
func Signup(context *gin.Context) {
	// TODO 这里需要进行判断用户名是否存在
	username := context.PostForm("username")
	password := context.PostForm("password")

	passwordByte, err := util.GeneratePassword(password)
	if err != nil {
		fmt.Println("加密出错")
	}

	timer := time.Now()

	user := model.User{}
	user.Username = username
	user.Password = string(passwordByte)
	user.Birthday = timer
	user.CreateTime = timer
	user.Icon = "https://pic4.zhimg.com/80/v2-82f6a065d5451c5cf3c51bc840e78ada_720w.jpg"
	user.Nickname = "未命名用户"

	state := user.Insert(user)
	result := model.Result{
		Code:    http.StatusOK,
		Message: "查询成功",
		Data:    user,
	}

	if state == 0 {
		return
	}
	context.JSON(http.StatusOK, result)
}

// 通过手机号注册
func SignupByPhone(context *gin.Context) {
	// TODO 这里需要进行判断手机号是否被注册
	phone := context.PostForm("phone")

	timer := time.Now()

	user := model.User{}
	user.Phone = phone
	//user.Birthday = timer.Format("2006-01-02")
	user.Birthday = timer
	user.CreateTime = timer
	user.Icon = "https://pic4.zhimg.com/80/v2-82f6a065d5451c5cf3c51bc840e78ada_720w.jpg"
	user.Nickname = "未命名用户"

	state := user.Insert(user)
	result := model.Result{
		Code:    http.StatusOK,
		Message: "查询成功",
		Data:    user,
	}

	if state == 0 {
		return
	}
	context.JSON(http.StatusOK, result)
}

func Login(context *gin.Context) {
	json := &LoginForm{}
	if context.BindJSON(&json) != nil {
		context.JSON(http.StatusBadRequest, model.Result{
			Code:    http.StatusBadRequest,
			Message: "请求参数为空",
			Data:    nil,
		})
		return
	}

	user := model.User{}
	member := user.FindByUsername(json.Username)

	if member.IsEmpty() {
		context.JSON(http.StatusNotFound, model.Result{
			Code:    http.StatusNotFound,
			Message: "用户名不存在",
			Data:    nil,
		})
		return
	}

	isOk, _ := util.ValidatePassword(json.Password, member.Password)
	if !isOk {
		context.JSON(http.StatusBadRequest, model.Result{
			Code:    http.StatusBadRequest,
			Message: "密码错误",
			Data:    nil,
		})
		return
	}

	context.JSON(http.StatusOK, member)
}
