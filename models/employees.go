package models

import "time"

type Employee struct {
	ID        int       `gorm:"serial;primary_key;autoIncrement"`
	FirstName string    `gorm:"type:text" json:"first_name"`
	LastName  string    `gorm:"type:text" json:"last_name"`
	Email     string    `gorm:"type:text" json:"email"`
	HireDate  time.Time `json:"hire_date,omitempty"`
}

type Emplpyees []Employee
