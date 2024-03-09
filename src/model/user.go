package model

import (
	"time"

	"gorm.io/gorm"
)

// User struct
type User struct {
	gorm.Model
	FirstName      string    `gorm:"not null; size:255" json:"firstname"`
	LastName       string    `gorm:"not null; size:255" json:"lastname"`
	Email          string    `gorm:"not null; size:255; unique" json:"email"`
	PhoneNumber    string    `gorm:"not null; size:255; unique" json:"phonenumber"`
	Password       string    `gorm:"not null; size:255" json:"password"`
	RegistrationId string    `gorm:"not null; size:255; unique" json:"registrationid"`
	DateOfBirth    string    `gorm:"not null; size:255" json:"dateofbirth"`
	CreatedAt      time.Time `gorm:"not null; size:255" json:"createdat"`
	UpdatedAt      time.Time `gorm:"not null; size:255" json:"updatedat"`
	DeletedAt      time.Time `gorm:"not null; size:255" json:"deletedat"`
}
