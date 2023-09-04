package apis

import (
	"context"
	"grpc_client_demo/client"
	"grpc_client_demo/notifications"

	"github.com/gin-gonic/gin"
)

func JsonError(err error) interface{} {
	v := make(map[string]interface{})
	v["error"] = err.Error()
	return v
}

func GetAllNotifications(c *gin.Context) {
	notifications, err := client.ServiceClient.GetAllNotifications(context.Background(), &notifications.Empty{})
	if err != nil || len(notifications.NotificationList) == 0 {
		c.JSON(500, JsonError(err))
		return
	}
	c.JSON(200, notifications.NotificationList)
}

func AddNotification(c *gin.Context) {
	title := c.Param("title")
	content := c.Param("content")
	notifications, err := client.ServiceClient.AddNotifications(context.Background(), &notifications.Notification{
		Title:   title,
		Content: content,
	})
	if err != nil {
		c.JSON(500, JsonError(err))
		return
	}
	c.JSON(200, notifications.NotificationList)
}

func RemoveNotification(c *gin.Context) {
	id := c.Param("id")
	_, err := client.ServiceClient.RemoveNotifications(context.Background(), &notifications.NotificationID{Id: id})
	if err != nil {
		c.JSON(500, JsonError(err))
		return
	}
	c.String(200, "ok")
}
