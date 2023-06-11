package models

import "time"

type Lesson struct {
	Name      string    `json:"name"`
	Teacher   string    `json:"teacher"`
	TimeStart time.Time `json:"time_start"`
	TimeEnd   time.Time `json:"time_end"`
}
