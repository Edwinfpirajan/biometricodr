package models

import (
	"time"
)

// type Employe struct {
// 	ID         int    `json:"id" gorm:"primary_key;auto_increment"`
// 	PinEmploye string `json:"pinEmploye" gorm:"FOREIGNKEY:PinEmploye" `
// 	FirstName  string `json:"first_name" `
// 	LastName   string `json:"last_name"`
// 	Company    string `json:"company"`
// 	Position   string `json:"position"`
// 	ScheduleId int    `json:"schedule"`
// 	/* 	Attendance Attendances `gorm:"type:bytes;serializer:gob"` */
// 	CreatedAt time.Time `json:"fechacreacion"`
// }

type Employe struct {
	ID         int    `json:"id" gorm:"primary_key;auto_increment"`
	PinEmploye string `json:"pinEmploye" gorm:"FOREIGNKEY:PinEmploye" `
	FirstName  string `json:"first_name" `
	LastName   string `json:"last_name"`
	Company    string `json:"company"`
	Position   string `json:"position"`
	ScheduleId int    `json:"schedule_id"`
	// Arrival    time.Time `json:"arrival"`
	// Departure  time.Time `json:"departure"`
	CreatedAt time.Time `json:"fechacreacion"`
}

type EmployeeWithSchedule struct {
	Employe
	Arrival   string `json:"arrival"`
	Departure string `json:"departure"`
}

// func (e EmployeeWithSchedule) TableName() string {
// 	return "employes"
// }

/* UpdatedAt  time.Time
DeletedAt  *time.Time */
