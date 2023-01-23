package entity

type Horary struct {
	Id_sch    int    `json:"id_sch" param:"id"`
	Arrival   string `json:"arrival"`
	Departure string `json:"departure"`
}
