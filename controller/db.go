package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"myapp/service"
	"net/http"
)

func Migrate(c *gin.Context) {
	err := service.Migrate()
	if err != nil {
		log.Println("迁移表结构：", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"result": false, "error": error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": true})
}
