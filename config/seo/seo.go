package seo

import (
	"github.com/qor/l10n"
	"github.com/qor/seo"
)

type MySEOSetting struct {
	seo.QorSEOSetting
	l10n.Locale
}

type SEOGlobalSetting struct {
	SiteName string
}

var SEOCollection *seo.Collection
