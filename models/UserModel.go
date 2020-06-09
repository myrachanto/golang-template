package models

import (
	"github.com/jinzhu/gorm"
	"time"
)
type User struct {
	FName string `gorm:"not null"`
	LName string `gorm:"not null"`
	UName string `gorm:"not null"`
	Phone string `gorm:"not null"`
	Address string `gorm:"not null"`
	Dob *time.Time  `gorm:"not null"`
	Picture string `gorm:"not null"`
	Email string `gorm:"not null;unique"`
	Password string `gorm:"not null"`
	gorm.Model
}
type Auth struct {
	User User `gorm:"foreignKey:UserID; not null"`
	UserID int `json:"userid"`
	Token string `gorm:"type:varchar(200); not null"`
	gorm.Model
}