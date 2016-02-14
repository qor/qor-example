package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type Phone struct {
	gorm.Model
	UserID uint
	Phone  string
}

func (phone Phone) Stringify() string {
	return fmt.Sprintf("%v", phone.Phone)
}
