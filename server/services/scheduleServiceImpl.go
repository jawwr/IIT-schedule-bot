package services

import (
	"server/models"
	"server/repository"
)

type ScheduleServiceImpl struct {
	repository *repository.MongoRepository
}

func NewScheduleService(repository *repository.MongoRepository) *ScheduleServiceImpl {
	return &ScheduleServiceImpl{
		repository: repository,
	}
}

func (s ScheduleServiceImpl) GetToday() *models.DaySchedule {
	return s.repository.GetSchedule()
}

func (s ScheduleServiceImpl) GetWeek() *models.WeekSchedule {
	number := 1
	return s.repository.GetWeekByNumber(number)
}

func (s ScheduleServiceImpl) GetLesson() *models.Lesson {
	return s.repository.GetLesson()
}

func (s ScheduleServiceImpl) GetWeekByNumber(number int) *models.WeekSchedule {
	return s.repository.GetWeekByNumber(number)
}
