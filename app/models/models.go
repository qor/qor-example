package models

import (
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/qor/qor/l10n"
	"github.com/qor/qor/media_library"
	"github.com/qor/qor/publish"
)

type Author struct {
	gorm.Model
	publish.Status
	l10n.Locale

	Name string
}

type Book struct {
	gorm.Model
	publish.Status
	l10n.Locale

	Title       string
	Synopsis    string
	ReleaseDate time.Time
	Authors     []*Author `gorm:"many2many:book_authors"`
	Price       float64
	CoverImage  media_library.FileSystem
	// later
	// CoverImages []ProductImage // product image has BookId => handles relation
}

// type ProductImage struct {
// 	gorm.Model
// 	publish.Status
// 	l10n.Locale

// 	BookId     uint
// 	CoverImage media_library.FileSystem
// }

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
		log.Println("EditableLocales() admin")
		return []string{l10n.Global, "ja-JP"}
		// return []string{l10n.Global}
	} else {
		//log.Println("EditableLocales() NOT admin")
		//return []string{l10n.Global, "ja-JP"}
		return []string{}
	}
}
