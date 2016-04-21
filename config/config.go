package config

import (
	"os"

	"github.com/jinzhu/configor"
	"github.com/qor/render"
)

var Config = struct {
	Port uint `default:"7000" env:"PORT"`
	DB   struct {
		Name     string `default:"qor_example"`
		Adapter  string `default:"mysql"`
		Host     string `default:"localhost"`
		User     string
		Password string
	}
}{}

var (
	Root = os.Getenv("GOPATH") + "/src/github.com/qor/qor-example"
	View *render.Render
)

func init() {
	cfgPath := os.Getenv("QOR_DBPATH")
	if cfgPath == "" {
		cfgPath = "config/database.yml"
	}

	if err := configor.Load(&Config, cfgPath); err != nil {
		panic(err)
	}

	View = render.New()
}
