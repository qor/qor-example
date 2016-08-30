package config

import (
	"os"

	"github.com/jinzhu/configor"
	"github.com/qor/render"
)

type SMTPConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Site     string
}

var Config = struct {
	Port uint `default:"7000" env:"PORT"`
	DB   struct {
		Name     string `default:"qor_example"`
		Adapter  string `default:"mysql"`
		User     string
		Password string
	}
	SMTP SMTPConfig
}{}

var (
	Root = os.Getenv("GOPATH") + "/src/github.com/qor/qor-example"
	View *render.Render
)

func init() {
	if err := configor.Load(&Config, "config/database.yml", "config/smtp.yml"); err != nil {
		panic(err)
	}

	View = render.New()
}

func (s SMTPConfig) HostWithPort() string {
	return s.Host + ":" + s.Port
}
