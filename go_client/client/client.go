package client

import (
	"grpc_client_demo/notifications"
	"log"

	"google.golang.org/grpc"
)

var Conn *grpc.ClientConn
var ServiceClient notifications.NotificationServiceClient

func InitGrpc() {
	var err error
	Conn, err = grpc.Dial(":30043", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	ServiceClient = notifications.NewNotificationServiceClient(Conn)
}
