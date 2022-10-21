package routes

import (
	"context"
	"go-learn/src/model"
	"go-learn/src/utils"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetProductsType struct {
	Prize      int
	Name       string
	Desc       string
	CreateTime string
	Likes      int
}

type RoutesType struct {
	Name string
	Path string
}

// 新增商品
func PostProducts(c *gin.Context) {
	// params
	desc := c.PostForm("desc")
	name := c.PostForm("name")
	prize := c.DefaultPostForm("prize", "0")

	if name == "" {
		c.SecureJSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "fail",
			"data": "name 不能为空",
		})
		return
	}

	var (
		products model.ProductInfo
	)

	collection := *utils.MongoDb()

	products.Desc = desc
	products.Prize = prize
	products.ProductName = name

	res, _ := collection.InsertOne(context.TODO(), products)

	c.SecureJSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "success",
		"data": res,
	})
}

// 获取商品信息
func GetProducts(c *gin.Context) {
	sale := [5]string{"商品1", "商品2", "商品3", "商品4", "商品5"}

	slice := []GetProductsType{}

	for _, v := range sale {
		var products GetProductsType
		products.Name = v
		products.Likes = rand.Intn(1000)
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

// 获取商品信息
func GetNavs(c *gin.Context) {
	sale := [5]string{"热销", "二手", "精品", "推荐"}

	slice := []RoutesType{}

	for _, v := range sale {
		var products RoutesType
		products.Name = v
		products.Path = "https://www.baidu.com"
		slice = append(slice, products)
	}

	c.SecureJSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "success",
		"data": slice,
	})
}

// 获取点赞数量
func GetLikes(c *gin.Context) {
	c.SecureJSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "success",
		"data": true,
	})
}
