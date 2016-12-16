package models

import (
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/qor/l10n"
	"github.com/qor/sorting"
	"github.com/qor/validations"
)

type Color struct {
	gorm.Model
	l10n.Locale
	sorting.Sorting
	Name string
	Code string `l10n:"sync"`
}

func (color Color) Validate(db *gorm.DB) {
	if strings.TrimSpace(color.Name) == "" {
		db.AddError(validations.NewError(color, "Name", "Name can not be empty"))
	}

	if strings.TrimSpace(color.Code) == "" {
		db.AddError(validations.NewError(color, "Code", "Code can not be empty"))
	}
}
