package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type Address struct {
	gorm.Model
	UserID      uint
	ContactName string
	Phone       string
	City        string
	Address1    string
	Address2    string
}

func (address Address) Stringify() string {
	return fmt.Sprintf("%s, %s", address.Address1, address.Address2)
}
