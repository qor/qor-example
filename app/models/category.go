package models

import (
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/qor/l10n"
	"github.com/qor/slug"
	"github.com/qor/sorting"
	"github.com/qor/validations"
)

type Category struct {
	gorm.Model
	l10n.Locale
	sorting.Sorting
	Name         string
	NameWithSlug slug.Slug

	Categories []Category
	CategoryID uint
}

func (category Category) Validate(db *gorm.DB) {
	if strings.TrimSpace(category.Name) == "" {
		db.AddError(validations.NewError(category, "Name", "Name can not be empty"))
	}
}

func (category Category) DefaultPath() string {
	defaultPath := "/"
	if len(category.Name) > 0 {
		defaultPath = fmt.Sprintf("/category/%s", category.NameWithSlug.Slug)
	}
	return defaultPath
}
