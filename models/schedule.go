package models

import "time"

type Horary struct {
	Id_sch    int       `json:"id_sch" gorm:"primary_key;auto_increment"`
	Arrival   time.Time `json:"arrival"`
	Departure time.Time `json:"departure"`
}
