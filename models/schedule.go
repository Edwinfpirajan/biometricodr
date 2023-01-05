package models

type Horary struct {
	ID                int    `json:"id" gorm:"primary_key;auto_increment"`
	ArrivalSchedule   string `json:"arrival"`
	DepartureSchedule string `json:"departure"`
}
