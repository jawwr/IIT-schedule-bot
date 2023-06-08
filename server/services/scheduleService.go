package services

import "server/models"

type ScheduleService interface {
	GetToday() *models.DaySchedule
	GetWeek() *models.WeekSchedule
	GetLesson() *models.Lesson
	GetWeekByNumber(number int) *models.WeekSchedule
}
