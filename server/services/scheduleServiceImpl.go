package services

import (
	models "server/models"
)

type ScheduleServiceImpl struct {
}

func (s ScheduleServiceImpl) GetToday() *models.DaySchedule {
	return &models.DaySchedule{
		Name: "Среда",
		Lessons: []*models.Lesson{
			{
				Name:      "что-то 1",
				Teacher:   "smth",
				TimeStart: "11:20",
				TimeEnd:   "12:50",
			},
			{
				Name:      "что-то 2",
				Teacher:   "smth",
				TimeStart: "13:15",
				TimeEnd:   "14:45",
			},
			{
				Name:      "что-то 3",
				Teacher:   "smth",
				TimeStart: "15:00",
				TimeEnd:   "16:30",
			},
		},
	}
}
