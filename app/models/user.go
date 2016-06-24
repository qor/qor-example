package models

import (
	"github.com/asaskevich/govalidator"
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	gorm.Model
	Email     string `valid:"required,email,uniqEmail~Email already be token"`
	Password  string `valid:"password~Password can't be blank"`
	Name      string `valid:"required"`
	Gender    string
	Role      string
	Addresses []Address

	// Confirm
	ConfirmToken string
	Confirmed    bool

	// Recover
	RecoverToken       string
	RecoverTokenExpiry time.Time
}

func (user *User) Validate(db *gorm.DB) {
	govalidator.CustomTypeTagMap.Set("uniqEmail", govalidator.CustomTypeValidator(func(email interface{}, context interface{}) bool {
		var user User
		db.First(&user, "email = ?", email)
		if user.ID == 0 || user.ID == context.(User).ID {
			return true
		}
		return false
	}))

	govalidator.CustomTypeTagMap.Set("password", govalidator.CustomTypeValidator(func(email interface{}, context interface{}) bool {
		if context.(User).ID == 0 && context.(User).Password == "" {
			return false
		}
		return true
	}))
}

func (user User) DisplayName() string {
	return user.Email
}

func (user User) AvailableLocales() []string {
	return []string{"en-US", "zh-CN"}
}
