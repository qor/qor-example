package config

import (
	"github.com/jinzhu/configor"
	"github.com/qor/qor/i18n"
)

var Config = struct {
	Port uint `default:"7000" env:"PORT"`
	DB   struct {
		Name     string `default:"qor_example"`
		Adapter  string `default:"mysql"`
		User     string
		Password string
		Host     string `default:"localhost"`
		Port     uint   `default:"3306"`
		Debug    bool   `default:"false"`
	}
	Redis struct {
		Host     string `default:"localhost"`
		Port     uint   `default:"6379"`
		Protocol string `default:"tcp"`
		Password string
	}
	Session struct {
		Name    string `default:"sessionid"`
		Adapter string `default:"cookie"`
	}
	I18n   *i18n.I18n
	Secret string `default:"secret"`
	Limit  int    `default:"5"`
}{}

func init() {
	if err := configor.Load(&Config, "config/database.yml"); err != nil {
		panic(err)
	}
}
