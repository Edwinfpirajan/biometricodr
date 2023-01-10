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
	FirstName    string     `json:"first_name" `
	LastName     string     `json:"last_name"`
	Arrival      *time.Time `json:"arrival"`
	BreakIn      *time.Time `json:"breakIn"`
	BreakOut     *time.Time `json:"breakOut"`
	Departure    *time.Time `json:"departure"`
	Photo        string     `json:"photo"`
}

// type FormData struct {
// 	PinEmpleado string
// 	Estado      string
// 	Foto        []byte
// }
