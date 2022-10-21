package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mgoCli *mongo.Client

type Config struct {
	Mongo string `json:"Mongo"`
}

func initEngine() {
	var err error

	// ===== start ====读取配置
	configFile, err := os.Open("config.json")

	if err != nil {
		fmt.Println(err)
	}

	// 关闭文件
	defer configFile.Close()

	byteValue, _ := ioutil.ReadAll(configFile)

	var config Config

	json.Unmarshal([]byte(byteValue), &config)
	// ===== end ====读取配置
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	clientOptions := options.Client().ApplyURI(config.Mongo)

	// 连接到MongoDB
	mgoCli, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// 检查连接
	err = mgoCli.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
}
func GetMgoCli() *mongo.Client {
	if mgoCli == nil {
		initEngine()
	}
	return mgoCli
}

func MongoDb() *mongo.Collection {
	var (
		client     = GetMgoCli()
		db         *mongo.Database
		collection *mongo.Collection
	)

	//2.选择数据库 my_db
	db = client.Database("mizzle")

	collection = db.Collection("dev")

	return collection
}
