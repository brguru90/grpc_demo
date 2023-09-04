package main

import (
	"grpc_demo/apis"
	"grpc_demo/notifications"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":30043")
	if err != nil {
		panic(err)
	}

	s := apis.Server{}
	grpc_server := grpc.NewServer()
	notifications.RegisterNotificationServiceServer(grpc_server, &s)

	if err := grpc_server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %s", err)
	}

}
