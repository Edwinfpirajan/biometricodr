package models

import "time"

// type Asistencia struct {
// 	// RegistroID  int64     `json:"id" gorm:"primary_key;auto_increment"`
// 	EmpleadoPIN string `json:"pin"`
// 	Estado      string `json:"estado"`
// 	Image       []byte `json:"foto"`

// 	CreatedAt time.Time `json:"hora"`
// }

type Attendances struct {
	ID           int        `json:"id" gorm:"primary_key;auto_increment"`
	PinEmployeFK string     `json:"pinEmploye"`
	Arrival      *time.Time `json:"arrival"`
	BreakIn      *time.Time `json:"breakIn"`
	BreakOut     *time.Time `json:"breakOut"`
	Departure    *time.Time `json:"departure"`
	CreatedAt    time.Time  `json:"date"`
	Photo        string     `json:"photo"`
}

type GetAllAttendances struct {
	EmployeeWithSchedule
	Attendances
	// CreatedAt time.Time  `json:"date"`
	// Arrival   string     `json:"arrival"`
	// BreakIn   *time.Time `json:"breakIn"`
	// BreakOut  *time.Time `json:"breakOut"`
	// Departure string     `json:"departure"`
	// Photo     string     `json:"photo"`
}

// type FormData struct {
// 	PinEmpleado string
// 	Estado      string
// 	Foto        []byte
// }
