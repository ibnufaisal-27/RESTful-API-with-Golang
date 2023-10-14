package models

import "time"

const (
	DateLayoutISO = "02/01/2006"
)

type Employee struct {
	ID        int       `gorm:"serial;primary_key;autoIncrement" json:"id"`
	FirstName string    `gorm:"type:text" json:"first_name"`
	LastName  string    `gorm:"type:text" json:"last_name"`
	Email     string    `gorm:"type:text" json:"email"`
	HireDate  time.Time `json:"hire_date,omitempty"`
}

type PayloadEmployee struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required"`
	HireDate  string `json:"hire_date" binding:"required"`
}
