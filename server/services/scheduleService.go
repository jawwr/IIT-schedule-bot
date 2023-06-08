package services

import "server/models"

type ScheduleService interface {
	GetToday() *models.DaySchedule
}
