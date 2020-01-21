package main

import (
	"awesomeProject6/common"
	"awesomeProject6/controller"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	common.InitMysqlPool()
	r := gin.Default()
	r.POST("/add", controller.TaskAdd)
	r.Run(":8088") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}