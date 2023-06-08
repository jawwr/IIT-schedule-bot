package services

import (
	"math/rand"
	"server/models"
	"strconv"
)

type ScheduleServiceImpl struct {
}

func (s ScheduleServiceImpl) GetToday() *models.DaySchedule {
	return &models.DaySchedule{
		Name: "Среда",
		Lessons: []*models.Lesson{
			s.GetLesson(),
			s.GetLesson(),
			s.GetLesson(),
		},
	}
}

func (s ScheduleServiceImpl) GetWeek() *models.WeekSchedule {
	return &models.WeekSchedule{
		Name: "Неделя",
		Days: []*models.DaySchedule{
			s.GetToday(),
			s.GetToday(),
		},
	}
}

func (s ScheduleServiceImpl) GetLesson() *models.Lesson {
	number := rand.Intn(10)
	return &models.Lesson{
		Name:      "что-то " + strconv.Itoa(number),
		Teacher:   "smth",
		TimeStart: "15:00",
		TimeEnd:   "16:30",
	}
}

func (s ScheduleServiceImpl) GetWeekByNumber(number int) *models.WeekSchedule {
	week := s.GetWeek()
	week.Name += " " + strconv.Itoa(number)
	return week
}
