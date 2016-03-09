package models

import (
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/qor/l10n"
	"github.com/qor/publish"
	"github.com/qor/validations"
	"github.com/qor/sorting"
)

type Category struct {
	gorm.Model
	l10n.Locale
	publish.Status
	sorting.Sorting
	Name string
}

func (category Category) Validate(db *gorm.DB) {
	if strings.TrimSpace(category.Name) == "" {
		db.AddError(validations.NewError(category, "Name", "Name can not be empty"))
	}
}
