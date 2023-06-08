package models

type WeekSchedule struct {
	Days []*DaySchedule `json:"days"`
}
