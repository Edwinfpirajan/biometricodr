package entity

type Employe struct {
	PinEmploye string `json:"pinEmploye" gorm:"FOREIGNKEY:PinEmploye" `
	FirstName  string `json:"first_name" `
	LastName   string `json:"last_name"`
	Company    string `json:"company"`
	Position   string `json:"position"`
	ScheduleId int    `json:"schedule_id"`
}
