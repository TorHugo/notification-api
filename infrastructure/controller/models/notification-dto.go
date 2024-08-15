package models

import "notification-api/domain/model"

type NotificationDTO struct {
	Contact    string            `json:"contact" binding:"required"`
	Subject    string            `json:"subject"`
	Template   string            `json:"template" binding:"required"`
	Parameters []model.Parameter `json:"parameters"`
}
