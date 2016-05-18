package auth

import (
	"github.com/justinas/nosurf"
	"github.com/qor/i18n/inline_edit"
	"github.com/qor/qor-example/config"
	"github.com/qor/qor-example/config/i18n"
	"gopkg.in/authboss.v0"
	_ "gopkg.in/authboss.v0/auth"
	_ "gopkg.in/authboss.v0/register"
	"html/template"
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
	Auth.ViewsPath = "app/views/auth"
	Auth.LayoutPath = config.Root + "/app/views/layouts/application.tmpl"
	Auth.LayoutFuncMaker = layoutFunc
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

func layoutFunc(w http.ResponseWriter, r *http.Request) template.FuncMap {
	funcsMap := template.FuncMap{
		"render": func(s interface{}) string { return "" },
	}
	for k, v := range inline_edit.FuncMap(i18n.I18n, "en-US", false) {
		funcsMap[k] = v
	}
	return funcsMap
}
