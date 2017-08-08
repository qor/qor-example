package auth

import (
	"github.com/qor/auth"
	"github.com/qor/auth_themes/clean"
	"github.com/qor/qor-example/app/models"
	"github.com/qor/qor-example/config"
	"github.com/qor/qor-example/db"
)

// Auth initialize auth
var Auth = clean.New(&auth.Config{
	DB:         db.DB,
	Render:     config.View,
	Mailer:     config.Mailer,
	UserModel:  models.User{},
	Redirector: auth.Redirector{config.RedirectBack},
})

// var Auth = auth.New(&auth.Config{
// 	DB:        db.DB,
// 	Render:    config.View,
// 	Mailer:    config.Mailer,
// 	UserModel: models.User{},
// })

// func init() {
// 	Auth.RegisterProvider(password.New(&password.Config{Confirmable: true}))
// 	Auth.RegisterProvider(phone.New())
// 	Auth.RegisterProvider(github.New(&config.Config.Github))
// 	Auth.RegisterProvider(google.New(&config.Config.Google))
// }
