package auth

import (
	"github.com/qor/auth"
	"github.com/qor/auth/providers/github"
	"github.com/qor/auth/providers/google"
	"github.com/qor/auth/providers/password"
	"github.com/qor/auth/providers/phone"
	"github.com/qor/qor-example/app/models"
	"github.com/qor/qor-example/config"
	"github.com/qor/qor-example/db"
)

// Auth initialize auth
var Auth = auth.New(&auth.Config{
	DB:        db.DB,
	Render:    config.View,
	Mailer:    config.Mailer,
	UserModel: models.User{},
})

func init() {
	Auth.RegisterProvider(password.New(&password.Config{Confirmable: true}))
	Auth.RegisterProvider(phone.New())
	Auth.RegisterProvider(github.New(&config.Config.Github))
	Auth.RegisterProvider(google.New(&config.Config.Google))
}
