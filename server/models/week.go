package models

type WeekSchedule struct {
	Name string         `json:"name"`
	Days []*DaySchedule `json:"days"`
}
