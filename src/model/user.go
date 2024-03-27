package model

import (
	"aqua-ratsnake/src/database"

	"gorm.io/gorm"
)

// User struct
type User struct {
	gorm.Model
	FirstName      string `gorm:"not null; size:255" json:"firstname"`
	LastName       string `gorm:"not null; size:255" json:"lastname"`
	Email          string `gorm:"not null; size:255; unique" json:"email"`
	PhoneNumber    string `gorm:"not null; size:255; unique" json:"phonenumber"`
	Password       string `gorm:"not null; size:255" json:"password"`
	RegistrationId string `gorm:"not null; size:255; unique" json:"registrationid"`
	DateOfBirth    string `gorm:"not null; size:255" json:"dateofbirth"`
}

// User Update struct
type UserUpdateInput struct {
	gorm.Model
	FirstName      string `gorm:"not null; size:255" json:"firstname"`
	LastName       string `gorm:"not null; size:255" json:"lastname"`
	Email          string `gorm:"not null; size:255; unique" json:"email"`
	PhoneNumber    string `gorm:"not null; size:255; unique" json:"phonenumber"`
	Password       string `gorm:"not null; size:255" json:"password"`
	DateOfBirth    string `gorm:"not null; size:255" json:"dateofbirth"`
}

// Save User
func (user *User) Save() (*User, error) {
	result := database.Database.Create(&user)
	if result.Error != nil {
		return &User{}, result.Error
	}
	return user, nil
}

