package models

type DaySchedule struct {
	Name    string    `json:"name"`
	Lessons []*Lesson `json:"lessons"`
}
