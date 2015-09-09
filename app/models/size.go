package models

import (
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/qor/qor/l10n"
	"github.com/qor/qor/publish"
	"github.com/qor/qor/sorting"
	"github.com/qor/qor/validations"
)

type Size struct {
	gorm.Model
	l10n.Locale
	publish.Status
	sorting.Sorting
	Name string
	Code string `l10n:"sync"`
}

func (size Size) Validate(db *gorm.DB) {
	if strings.Trim(size.Name, " ") == "" {
		db.AddError(validations.NewError(size, "Name", "Name can not be empty"))
	}

	if strings.Trim(size.Code, " ") == "" {
		db.AddError(validations.NewError(size, "Code", "Code can not be empty"))
	}
}
