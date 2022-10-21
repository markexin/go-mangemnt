package model

type ProductInfo struct {
	// 商品名称
	ProductName string `bson:"projectName"`
	// 商品价格
	Prize string `bson:"prize"`
	// 商品描述
	Desc string `bson:"desc"`
	//开始时间
	StartTime int64 `bson:"startTime"`
	//结束时间
	EndTime int64 `bson:"endTime"`
}
