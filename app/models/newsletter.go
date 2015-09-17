package models

import "github.com/jinzhu/gorm"

type Newsletter struct {
	gorm.Model
	UserID         uint
	User           User
	Email          string
	Status         string
	NewsletterType string
}
