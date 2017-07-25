package auth

import (
	"bytes"
	"fmt"
	"io"
	"strings"

	"github.com/qor/auth"
	"github.com/qor/auth_themes/clean"
	"github.com/qor/mailer"
	"github.com/qor/mailer/gomailer"
	"github.com/qor/qor-example/app/models"
	"github.com/qor/qor-example/config"
	"github.com/qor/qor-example/db"
	gomail "gopkg.in/gomail.v2"
)

// Auth initialize auth
var Auth = clean.New(&auth.Config{
	DB:        db.DB,
	Render:    config.View,
	UserModel: models.User{},
})

func init() {
	sender := gomail.SendFunc(gomail.SendFunc(func(from string, to []string, msg io.WriterTo) error {
		var buf bytes.Buffer
		fmt.Printf("From: %v\n", from)
		fmt.Printf("To: %v\n", strings.Join(to, ", "))
		msg.WriteTo(&buf)
		fmt.Println(buf.String())
		return nil
	}))

	Auth.Config.Mailer = mailer.New(&mailer.Config{
		Sender: gomailer.New(&gomailer.Config{Sender: sender}),
	})
}
