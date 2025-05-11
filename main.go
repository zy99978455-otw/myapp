package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"myapp/controller"
)

func main() {
	// 初始化gin
	r := gin.Default()

	// 数据库同步
	r.POST("/db/sync", controller.Migrate)

	user := r.Group("/user")
	{
		// 用户注册
		user.POST("/register", controller.Register)

		//用户登录
		user.POST("/login", controller.Login)

		// 根据用户名查询用户信息
		user.GET("/getByName", controller.GetUserByName)

		// 删除用户
		user.DELETE("/:id", controller.DeleteById)

	}

	// 启动服务器
	if err := r.Run(":18080"); err != nil {
		log.Fatal(err)
	}

}
