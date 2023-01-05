package entity

type Attendance struct {
	PinEmployeFK string `json:"pinEmploye"`
	State        string `json:"state"`
	Photo        string `json:"photo"`
}
