package models

import "time"

type Attendances struct {
	ID           int        `json:"id" gorm:"primary_key;auto_increment"`
	PinEmployeFK string     `json:"pinEmploye"`
	Arrival      *time.Time `json:"arrival"`
	BreakInit    *time.Time `json:"breakInit"`
	BreakEnd     *time.Time `json:"breakEnd"`
	BreakInitTwo *time.Time `json:"breakInitTwo"`
	BreakEndTwo  *time.Time `json:"breakEndTwo"`
	BreakIn      *time.Time `json:"breakIn"`
	BreakOut     *time.Time `json:"breakOut"`
	Departure    *time.Time `json:"departure"`
	CreatedAt    time.Time  `json:"date"`
	Photo        string     `json:"photo"`
}

type GetAllAttendances struct {
	EmployeeWithSchedule
	Attendances
}
