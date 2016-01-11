package config

import (
	"github.com/jinzhu/configor"
	"github.com/qor/i18n"
)

var Config = struct {
	Port uint `default:"7000" env:"PORT"`
	DB   struct {
		Name     string `default:"qor_example"`
		Adapter  string `default:"mysql"`
		User     string
		Password string
	}
	I18n *i18n.I18n
}{}

func init() {
	if err := configor.Load(&Config, "config/database.yml"); err != nil {
		panic(err)
	}
}
