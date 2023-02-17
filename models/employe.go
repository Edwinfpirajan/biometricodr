package models

import "time"

type Employe struct {
	ID         int       `json:"id" gorm:"primary_key;auto_increment"`
	PinEmploye string    `json:"pinEmploye" gorm:"FOREIGNKEY:PinEmploye" `
	FirstName  string    `json:"first_name" `
	LastName   string    `json:"last_name"`
	Company    string    `json:"company"`
	Position   string    `json:"position"`
	ScheduleId int       `json:"schedule_id"`
	CreatedAt  time.Time `json:"fechacreacion"`
}

type EmployeeWithSchedule struct {
	Employe
	Horary
}
