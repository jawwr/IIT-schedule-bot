package models

type Group struct {
	Id         int64         `json:"_id"`
	Name       string        `json:"name"`
	FirstWeek  *WeekSchedule `json:"first_week"`
	SecondWeek *WeekSchedule `json:"second_week"`
}
