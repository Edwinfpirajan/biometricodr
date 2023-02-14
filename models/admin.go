package models

type Admin struct {
	ID       int    `json:"id" gorm:"primary_key;auto_increment"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
