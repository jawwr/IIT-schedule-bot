package models

type Group struct {
	Name       string        `json:"name"`
	FirstWeek  *WeekSchedule `json:"first_week"`
	SecondWeek *WeekSchedule `json:"second_week"`
}
