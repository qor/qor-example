package config

import "github.com/jinzhu/configor"

var Config = struct {
	DB struct {
		Name     string `default:"qor-example"`
		Adapter  string `default:"mysql"`
		Username string
		Password string
	}
}{}

func init() {
	configor.Load(&Config, "config/database.yml")
}
