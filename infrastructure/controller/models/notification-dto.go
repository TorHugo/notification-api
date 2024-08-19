package models

import (
	"notification-api/domain"
)

type NotificationDTO struct {
	Contact    string             `json:"contact" binding:"required"`
	Subject    string             `json:"subject"`
	Template   string             `json:"template" binding:"required"`
	Parameters []domain.Parameter `json:"parameters"`
}
