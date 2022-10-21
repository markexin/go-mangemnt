package main

import (
	"go-learn/src/routes"
	"net/http"

	"github.com/gin-gonic/gin"
)

var PREFIX string = "/api"

func main() {
	// 1.创建路由
	r := gin.Default()
	// 2.绑定路由规则，执行的函数
	// gin.Context，封装了request和response
	r.POST(PREFIX+"/login", routes.Login)
	r.POST(PREFIX+"/add_products", routes.PostProducts)
	r.GET(PREFIX+"/products", routes.GetProducts)
	r.GET(PREFIX+"/navs", routes.GetNavs)
	r.GET(PREFIX+"/likes", routes.GetLikes)
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "hi welcome~~~",
		})
	})
	// 3.监听端口，默认在8080
	// Run("里面不指定端口号默认为8080")
	r.Run(":8001")
}
