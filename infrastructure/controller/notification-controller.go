package controller

import "github.com/gin-gonic/gin"

type NotificationController struct {
	// use-case
}

func NewNotificationController() NotificationController {
	return NotificationController{}
}

func (p *NotificationController) SendNotification(ctx *gin.Context) {
	ctx.JSON(200, "")
}
