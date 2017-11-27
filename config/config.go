package config

import (
	"html/template"
	"os"

	"github.com/jinzhu/configor"
	"github.com/microcosm-cc/bluemonday"
	"github.com/qor/auth/providers/facebook"
	"github.com/qor/auth/providers/github"
	"github.com/qor/auth/providers/google"
	"github.com/qor/auth/providers/twitter"
	"github.com/qor/location"
	"github.com/qor/mailer"
	"github.com/qor/mailer/logger"
	"github.com/qor/media/oss"
	"github.com/qor/oss/s3"
	"github.com/qor/redirect_back"
	"github.com/qor/render"
	"github.com/qor/session/manager"
)

type SMTPConfig struct {
	Host     string
	Port     string
	User     string
	Password string
}

var Config = struct {
	Port uint `default:"7000" env:"PORT"`
	DB   struct {
		Name     string `env:"DBName" default:"qor_example"`
		Adapter  string `env:"DBAdapter" default:"mysql"`
		Host     string `env:"DBHost" default:"localhost"`
		Port     string `env:"DBPort" default:"3306"`
		User     string `env:"DBUser"`
		Password string `env:"DBPassword"`
	}
	S3 struct {
		AccessKeyID     string `env:"AWS_ACCESS_KEY_ID"`
		SecretAccessKey string `env:"AWS_SECRET_ACCESS_KEY"`
		Region          string `env:"AWS_Region"`
		S3Bucket        string `env:"AWS_Bucket"`
	}
	SMTP         SMTPConfig
	Github       github.Config
	Google       google.Config
	Facebook     facebook.Config
	Twitter      twitter.Config
	GoogleAPIKey string `env:"GoogleAPIKey"`
	BaiduAPIKey  string `env:"BaiduAPIKey"`
}{}

var (
	Root         = os.Getenv("GOPATH") + "/src/github.com/qor/qor-example"
	View         *render.Render
	Mailer       *mailer.Mailer
	RedirectBack = redirect_back.New(&redirect_back.Config{
		SessionManager:  manager.SessionManager,
		IgnoredPrefixes: []string{"/auth"},
	})
)

func init() {
	if err := configor.Load(&Config, "config/database.yml", "config/smtp.yml", "config/application.yml"); err != nil {
		panic(err)
	}

	location.GoogleAPIKey = Config.GoogleAPIKey
	location.BaiduAPIKey = Config.BaiduAPIKey

	View = render.New(nil)

	if Config.S3.AccessKeyID != "" {
		oss.Storage = s3.New(&s3.Config{
			AccessID:  Config.S3.AccessKeyID,
			AccessKey: Config.S3.SecretAccessKey,
			Region:    Config.S3.Region,
			Bucket:    Config.S3.S3Bucket,
		})
	}

	htmlSanitizer := bluemonday.UGCPolicy()
	View.RegisterFuncMap("raw", func(str string) template.HTML {
		return template.HTML(htmlSanitizer.Sanitize(str))
	})

	// dialer := gomail.NewDialer(Config.SMTP.Host, Config.SMTP.Port, Config.SMTP.User, Config.SMTP.Password)
	// sender, err := dialer.Dial()

	// Mailer = mailer.New(&mailer.Config{
	// 	Sender: gomailer.New(&gomailer.Config{Sender: sender}),
	// })
	Mailer = mailer.New(&mailer.Config{
		Sender: logger.New(&logger.Config{}),
	})
}
