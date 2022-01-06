package main

import (
	"context"
	"fmt"
	"time"
	"github.com/gin-gonic/gin"
	"net/http"
	"go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
)

type userType struct {
	Name string
	Age  uint32
}
// mongodb参考
// https://github.com/mongodb/mongo-go-driver
// https://thiscalifornianlife.com/2021/01/05/golang-mongodb/
// https://qiita.com/daredeshow/items/3c247210e9b06c313c22
func main() {
	fmt.Println("application's running.")

	// Connect to MongoDB
	fmt.Println("connecting to MongoDB.")
	// 接続Auth設定参考
	// TODO: エラー発生中
	// https://qiita.com/dobusarai/items/1960a7a30e213092b7aa
	var cred options.Credential
	cred.AuthSource = "go_db"
	cred.Username = "root"
	cred.Password = "password!"

	// DB接続設定
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://go_db:27017").SetAuth(cred))
	if err != nil { fmt.Println("Server Not Found.") }
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil { fmt.Println("Connection Failed.") }
	defer client.Disconnect(ctx)

	// レコードの挿入
	insert := userType {
		Name: "Newly Created User.",
		Age: 27,
	}
	collect := client.Database("go_db").Collection("users")
	res, err := collect.InsertOne(context.Background(), insert)
	fmt.Println(res)
	fmt.Println(err)

	// Server Run
	fmt.Println("server's running...")
	engine := gin.Default()
	engine.GET("/", func (c *gin.Context) {
		c.JSON(http.StatusOK, gin.H {
			"message": "hello go.",
		})
	})
	engine.Run("0.0.0.0:3000")
}