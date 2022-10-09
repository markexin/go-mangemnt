package routes

import (
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetProductsType struct {
	Prize      int
	Name       string
	Desc       string
	CreateTime string
}

// 获取商品信息
func GetProducts(c *gin.Context) {
	sale := [5]string{"商品1", "商品2", "商品3", "商品4", "商品5"}

	slice := []GetProductsType{}

	for _, v := range sale {
		var products GetProductsType
		products.Name = v
		products.Prize = rand.Intn(100)
		products.Desc = "测试商品"
		slice = append(slice, products)
	}

	c.SecureJSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "success",
		"data": slice,
	})
}
