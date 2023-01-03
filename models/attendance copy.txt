package models

import "time"

// type Asistencia struct {
// 	// RegistroID  int64     `json:"id" gorm:"primary_key;auto_increment"`
// 	EmpleadoPIN string `json:"pin"`
// 	Estado      string `json:"estado"`
// 	Image       []byte `json:"foto"`

// 	CreatedAt time.Time `json:"hora"`
// }

type Attendance struct {
	ID         uint      `gorm:"primary_key"`
	PinEmploye string    `json:"pinEmploye"`
	State      string    `json:"state"`
	Photo      string    `json:"photo"`
	CreatedAt  time.Time `json:"time"`
}

// type FormData struct {
// 	PinEmpleado string
// 	Estado      string
// 	Foto        []byte
// }
