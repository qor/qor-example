package seo

import (
	"github.com/qor/l10n"
	"github.com/qor/seo"
)

type MySeoSetting struct {
	seo.QorSeoSetting
	l10n.Locale
}

type SeoGlobalSetting struct {
	SiteName string
}

var SeoCollection *seo.Collection
