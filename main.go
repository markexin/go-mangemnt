package main

import (
	"go-learn/src/routes"

	"github.com/gin-gonic/gin"
)

var PREFIX string = "/api"

func main() {

	// 1.创建路由
	r := gin.Default()
	// 2.绑定路由规则，执行的函数
	// gin.Context，封装了request和response
	r.POST(PREFIX+"/login", routes.Login)
	r.GET(PREFIX+"/products", routes.GetProducts)
	r.GET(PREFIX+"/navs", routes.GetNavs)
	r.GET(PREFIX+"/likes", routes.GetLikes)
	// 3.监听端口，默认在8080
	// Run("里面不指定端口号默认为8080")
	r.Run(":8001")
}
