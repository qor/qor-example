package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/qor/qor/l10n"
	"github.com/qor/qor/publish"
)

type Author struct {
	gorm.Model
	publish.Status
	l10n.Locale

	Name string
}

type User struct {
	gorm.Model

	Name string
	Role string
}

func (u User) DisplayName() string {
	return u.Name
}

func (User) ViewableLocales() []string {
	return []string{l10n.Global, "ja-JP"}
}

func (user User) EditableLocales() []string {
	if user.Role == "admin" {
		return []string{l10n.Global, "ja-JP"}
	}
	return []string{}
}
