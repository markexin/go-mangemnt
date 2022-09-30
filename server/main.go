package main

import (
	"server/src/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/login", controllers.Login())
	r.Run(":8001")
}
