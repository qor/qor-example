package config

import "github.com/jinzhu/configor"

var Config = struct {
	Port uint `default:"7000"`
	DB   struct {
		Name     string `default:"qor-example"`
		Adapter  string `default:"mysql"`
		Username string
		Password string
	}
}{}

func init() {
	if err := configor.Load(&Config, "config/database.yml"); err != nil {
		panic(err)
	}
}
