package auth

import (
	"html/template"
	"net/http"
	"net/smtp"
	"os"
	"regexp"

	"github.com/gorilla/csrf"
	"github.com/qor/i18n/inline_edit"
	"github.com/qor/qor-example/config"
	"github.com/qor/qor-example/config/i18n"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/authboss.v0"
	_ "gopkg.in/authboss.v0/auth"
	_ "gopkg.in/authboss.v0/confirm"
	_ "gopkg.in/authboss.v0/recover"
	_ "gopkg.in/authboss.v0/register"
)

var (
	Auth = authboss.New()
)

func init() {
	Auth.MountPath = "/auth"
	Auth.XSRFName = "gorilla.csrf.Token"
	Auth.XSRFMaker = func(_ http.ResponseWriter, r *http.Request) string {
		return csrf.Token(r)
	}
	Auth.CookieStoreMaker = NewCookieStorer
	Auth.SessionStoreMaker = NewSessionStorer
	Auth.BCryptCost = bcrypt.DefaultCost
	Auth.LogWriter = os.Stdout
	Auth.Storer = &AuthStorer{}
	Auth.ViewsPath = "app/views/auth"
	Auth.LayoutPath = config.Root + "/app/views/layouts/application.tmpl"
	Auth.LayoutFuncMaker = layoutFunc
	Auth.LayoutDataMaker = layoutData
	Auth.Mailer = authboss.SMTPMailer(config.Config.SMTP.HostWithPort(), smtp.PlainAuth("", config.Config.SMTP.User, config.Config.SMTP.Password, config.Config.SMTP.Host))
	Auth.EmailFrom = "Qor Example"
	Auth.RootURL = config.Config.SMTP.Site
	Auth.Policies = []authboss.Validator{
		authboss.Rules{
			FieldName:       "email",
			Required:        true,
			AllowWhitespace: false,
			MustMatch:       regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`),
			MatchError:      "Please input a valid email address",
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

func CurrentLocale(req *http.Request) string {
	locale := "en-US"
	if cookie, err := req.Cookie("locale"); err == nil {
		locale = cookie.Value
	}
	return locale
}

func layoutData(w http.ResponseWriter, r *http.Request) authboss.HTMLData {
	return authboss.HTMLData{
		"Result": authboss.HTMLData{
			"CurrentLocale": CurrentLocale(r),
		},
	}
}

func layoutFunc(w http.ResponseWriter, r *http.Request) template.FuncMap {
	funcsMap := template.FuncMap{
		"render": func(s interface{}) string { return "" },
	}
	for k, v := range inline_edit.FuncMap(i18n.I18n, CurrentLocale(r), false) {
		funcsMap[k] = v
	}
	return funcsMap
}
