package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Metod basic/web/password/token/
type Auth struct {
	User     string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Metod    string `form:"metod" json:"metod" binding:"required"`
}

type LogLogin struct {
	gorm.Model
	UserID    uint
	User      User
	ClietIp   string
	LoginedAt *time.Time
	InOut     string
	Device    string
}
