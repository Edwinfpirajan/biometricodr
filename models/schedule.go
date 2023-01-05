package models

type Horary struct {
	ID            int    `json:"id" gorm:"primary_key;auto_increment"`
	ArrivalTime   string `json:"arrival_time" `
	BreakIn       string `json:"break_in" `
	BreakOut      string `json:"break_out"`
	DepartureTime string `json:"departure_time"`
	// Attendance    Attendance `gorm:"type:bytes;serializer:gob"`
}
