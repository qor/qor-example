package models

import (
	"github.com/jinzhu/gorm"
	"github.com/qor/seo"
)

type Seo struct {
	gorm.Model
	SiteName    string
	DefaultPage seo.Setting
	ProductPage seo.Setting `seo:"Name,Code"`
}
