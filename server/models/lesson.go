package models

type Lesson struct {
	Name      string `json:"name"`
	Teacher   string `json:"teacher"`
	TimeStart string `json:"time_start"`
	TimeEnd   string `json:"time_end"`
}
