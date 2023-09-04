package main

import (
	"grpc_client_demo/apis"
	"grpc_client_demo/client"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	client.InitGrpc()
	defer client.Conn.Close()

	router := gin.Default()
	router.GET("/", apis.GetAllNotifications)
	router.GET("/add/:title/:content", apis.AddNotification)
	router.GET("/delete/:id", apis.RemoveNotification)

	http.ListenAndServe(":3000", router)
}
