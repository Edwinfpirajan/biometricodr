package entity

type Attendance struct {
	PinEmployeFK string `json:"pinEmploye"`
	State        string `json:"state"`
	Photo        string `json:"photo"`
	// CreatedAt    time.Time `json:"date"`
}

type ValidateHorary struct {
	PinEmployeFK string `json:"pinEmploye"`
	Date         string `json:"date"`
}
