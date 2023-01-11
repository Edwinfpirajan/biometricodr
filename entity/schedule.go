package entity

import "time"

type Horary struct {
	ID        int       `json:"id" gorm:"primary_key;auto_increment"`
	Arrival   time.Time `json:"arrival"`
	Departure time.Time `json:"departure"`
}
