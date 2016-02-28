package models

import (
	"github.com/jinzhu/gorm"
	"github.com/qor/media_library"
)

type User struct {
	gorm.Model
	Email          string `sql:"type:varchar(75)" json:"email"`
	Name           string `gorm:"column:name" sql:"type:varchar(30);unique_index" json:"username"`
	Password       string `sql:"type:varchar(128)" json:"-"`
	IsActive       bool   `sql:"default:false" gorm:"column:is_active" json:"active"`
	FirstName      string `sql:"type:varchar(30)" json:"first_name"`
	LastName       string `sql:"type:varchar(30)" json:"last_name"`
	OrganizationID uint
	Organization   Organization
	Gender         string
	Role           string
	Languages      []Language `gorm:"many2many:user_languages;"`
	Addresses      []Address
	Comment        string
	// Role      Role
	// Email     []Email
	// Phone     []Phone
	// Social    []Social
	// Role      string
	// Location  string
	Avatar media_library.FileSystem
}

// func (user User) TableName() string {
//  return "auth_user"
// }

func (user User) DisplayName() string {
	return user.Name
}

func (user User) AvailableLocales() []string {
	return []string{"en-US", "uk-UA", "ru-RU"}
}

// func (User) ViewableLocales() []string {
//   return []string{l10n.Global, "zh-CN", "JP", "EN", "DE"}
// }

// func (user User) EditableLocales() []string {
//   if user.role == "global_admin" {
//     return []string{l10n.Global, "zh-CN", "EN"}
//   } else {
//     return []string{"zh-CN", "EN"}
//   }
// }

type Language struct {
	gorm.Model
	Name string
	Code string
}

// User Role
type Role struct {
	gorm.Model
	Name string
}

// func Roles() (results []string) {
// 	roleVariations := []Role{}
// 	if err := db.DB.Debug().Find(&roleVariations).Error; err != nil {
// 		log.Fatalf("query Role (%v) failure, got err %v", roleVariations, err)
// 		return results
// 	}
// 	for _, role := range roleVariations {
// 		results = append(results, role.Name)
// 	}
// 	return results
// }
