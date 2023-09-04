package apis

import (
	"context"
	"grpc_demo/local_db"
	"grpc_demo/notifications"
	"time"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	notifications.UnimplementedNotificationServiceServer
}

func (s *Server) GetAllNotifications(ctx context.Context, call *notifications.Empty) (*notifications.Notifications, error) {
	if len(local_db.NotificationList) == 0 {
		return nil, status.Error(codes.NotFound, "Empty data")
	}
	return &notifications.Notifications{NotificationList: local_db.NotificationList}, nil
}

func (s *Server) AddNotifications(ctx context.Context, call *notifications.Notification) (*notifications.Notifications, error) {
	local_db.NotificationList = append(local_db.NotificationList, &notifications.Notification{
		Id:        uuid.New().String(),
		Title:     call.Title,
		Content:   call.Content,
		TimeStamp: uint64(time.Now().UnixMilli()),
	})
	return &notifications.Notifications{NotificationList: local_db.NotificationList}, nil
}

func (s *Server) RemoveNotifications(ctx context.Context, call *notifications.NotificationID) (*notifications.Empty, error) {
	find_index := 0
	for _, notification := range local_db.NotificationList {
		if notification.Id == call.Id {
			break
		}
		find_index += 1
	}
	if find_index == len(local_db.NotificationList) {
		return nil, status.Error(codes.NotFound, "Not found")
	}
	local_db.NotificationList = append(local_db.NotificationList[:find_index], local_db.NotificationList[find_index+1:]...)
	return &notifications.Empty{}, nil
}
