package models

import (
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/qor/l10n"
	"github.com/qor/qor/publish"
	"github.com/qor/validations"
	"github.com/qor/sorting"
)

type Units struct {
	gorm.Model
	l10n.Locale
	publish.Status
	sorting.Sorting
	Name string
	Code string `l10n:"sync"`
}

func (units Units) Validate(db *gorm.DB) {
	if strings.TrimSpace(units.Name) == "" {
		db.AddError(validations.NewError(units, "Name", "Name can not be empty"))
	}

	if strings.TrimSpace(units.Code) == "" {
		db.AddError(validations.NewError(units, "Code", "Code can not be empty"))
	}
}
