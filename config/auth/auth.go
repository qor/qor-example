package auth

import (
	"github.com/justinas/nosurf"
	"gopkg.in/authboss.v0"
	_ "gopkg.in/authboss.v0/auth"
	_ "gopkg.in/authboss.v0/register"
	"net/http"
	"os"
)

var (
	Auth = authboss.New()
)

func init() {
	Auth.MountPath = "/auth"
	Auth.XSRFName = "csrf_token"
	Auth.XSRFMaker = func(_ http.ResponseWriter, r *http.Request) string {
		return nosurf.Token(r)
	}
	Auth.CookieStoreMaker = NewCookieStorer
	Auth.SessionStoreMaker = NewSessionStorer
	Auth.LogWriter = os.Stdout
	Auth.Storer = &AuthStorer{}
	Auth.Policies = []authboss.Validator{
		authboss.Rules{
			FieldName:       "email",
			Required:        true,
			AllowWhitespace: false,
		},
		authboss.Rules{
			FieldName:       "password",
			Required:        true,
			MinLength:       4,
			MaxLength:       8,
			AllowWhitespace: false,
		},
	}

	if err := Auth.Init(); err != nil {
		panic(err)
	}
}
