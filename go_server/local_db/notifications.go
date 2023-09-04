package local_db

import (
	"grpc_demo/notifications"
	"time"

	"github.com/google/uuid"
)

var NotificationList []*notifications.Notification = []*notifications.Notification{
	{
		Id:        uuid.New().String(),
		Title:     "test title",
		Content:   "test content",
		TimeStamp: uint64(time.Now().UnixMilli()),
	},
}
