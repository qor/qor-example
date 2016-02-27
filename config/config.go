package config

import (
	"os"

	"github.com/jinzhu/configor"
	"github.com/qor/i18n"
)

var Config = struct {
	SiteName string `default:"Qor DEMO"`
	Env      string `env:"ENV" default:"local"`
	Port     uint   `default:"7000" env:"PORT"`
	DB       struct {
		Name     string `default:"qor_example"`
		Adapter  string `default:"mysql"`
		User     string
		Password string
		Host     string `default:"localhost"`
		Port     uint   `default:"3306"`
		Debug    bool   `default:"false"`
	}
	Log struct {
		FileName string
		Maxdays  int `default:"30"`
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
	Locale string `default:"en-US"`
	Secret string `default:"secret"`
	Limit  int    `default:"5"`
}{}

var (
	Root       = os.Getenv("GOPATH") + "/src/github.com/qor/qor-example"
	FileConfig = "config/database.yml"
)

// Set environment variable config path -> export QORCONFIG=/etc/qor/production.yml
func init() {
	if file := os.Getenv("QORCONFIG"); len(file) > 0 {
		FileConfig = file
	}
	if rootPath := os.Getenv("QORROOT"); len(rootPath) > 0 {
		Root = rootPath
	}
	if err := configor.Load(&Config, FileConfig); err != nil {
		panic(err)
	}
}
