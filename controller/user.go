package controller

import (
	"log"
	"myapp/model"
	"myapp/service"
	"myapp/vo"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

func Register(c *gin.Context) {
	var req vo.UserReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"解析数据失败：": err.Error()})
		return
	}

	var user model.User
	copier.Copy(&user, &req)

	err := service.ValidateUser(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"数据校验失败：": err.Error()})
		return
	}

	err = service.RegisterUser(&user)

	if err != nil {
		log.Println("注册失败：", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error:": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"userId": user.ID})
}

func Login(c *gin.Context) {
	var req vo.UserReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"解析数据失败：": err.Error()})
		return
	}

	var user model.User
	copier.Copy(&user, &req)

	err := service.ValidateLogin(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"数据校验失败：": err.Error()})
		return
	}

	loginResult, err := service.Login(&user)

	c.JSON(http.StatusOK, gin.H{"result": loginResult, "error": err})
}

func GetUserByName(c *gin.Context) {
	username := c.Query("name")

	if len(username) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "用户名不能为空"})
		return
	}

	user, err := service.GetUserByName(username)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error:": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func DeleteById(c *gin.Context) {
	p := c.Param("id")
	id, err := strconv.Atoi(p)
	if err != nil {
		log.Println("id转为数字失败：", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = service.DeleteById(id)
	if err != nil {
		log.Println("删除用户失败：", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": true})
}
