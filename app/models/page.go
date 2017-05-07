package models

import (
	"github.com/jinzhu/gorm"
	"github.com/qor/publish2"
	"github.com/qor/slug"
	"github.com/qor/sorting"
	"github.com/qor/widget"
)

type Page struct {
	gorm.Model

	Title         string
	TitleWithSlug slug.Slug

	QorWidgetSettings       []widget.QorWidgetSetting `gorm:"many2many:page_qor_widget_settings;ForeignKey:id;AssociationForeignKey:name"`
	QorWidgetSettingsSorter sorting.SortableCollection

	publish2.Version
	publish2.Schedule
	publish2.Visible
}
