package models

import (
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/qor/qor/l10n"
	"github.com/qor/qor/publish"
	"github.com/qor/qor/sorting"
	"github.com/qor/qor/validations"
)

type Color struct {
	gorm.Model
	l10n.Locale
	publish.Status
	sorting.Sorting
	Name string
	Code string `l10n:"sync"`
}

func (color Color) Validate(db *gorm.DB) {
	if strings.Trim(color.Name, " ") == "" {
		db.AddError(validations.NewError(color, "Name", "Name can not be empty"))
	}

	if strings.Trim(color.Code, " ") == "" {
		db.AddError(validations.NewError(color, "Code", "Code can not be empty"))
	}
}
