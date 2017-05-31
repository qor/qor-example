package auth

import (
	"github.com/qor/auth"
	"github.com/qor/auth/database"
	"github.com/qor/auth/oauth/twitter"
	"github.com/qor/auth/phone"
)

var Auth = auth.New(nil)

func init() {
	Auth.RegisterProvider(database.New())
	Auth.RegisterProvider(phone.New())
	Auth.RegisterProvider(twitter.New())
}
