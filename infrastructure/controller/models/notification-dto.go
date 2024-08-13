package models

import "notification-api/domain/model"

type NotificationDTO struct {
	Contact    string            `json:"contact" binding:"required,email"`
	Subject    string            `json:"subject" binding:"required"`
	Template   string            `json:"template" binding:"required"`
	Parameters []model.Parameter `json:"parameters"`
}
