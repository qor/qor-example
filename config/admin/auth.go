package admin

import (
	"fmt"
	"log"

	"github.com/cryptix/go/http/auth"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"github.com/qor/admin"
	"github.com/qor/qor"
	"golang.org/x/crypto/bcrypt"

	"github.com/qor/qor-example/app/models"
	"github.com/qor/qor-example/db"
)

const (
	sessionName = "qor_example"
	authLanding = "/admin" // redirect location after login
	authLogout  = "/login" // redirect location after logout
)

type Auth struct {
	store   sessions.Store
	Handler *auth.Handler
}

func (a *Auth) Init() error {
	a.store = &sessions.CookieStore{
		Codecs: securecookie.CodecsFromPairs(
			[]byte("verysecretverysecretverysecret!!"), // TODO: gen and persist or panic
			[]byte("verysecretverysecretverysecret@@"),
		),
		Options: &sessions.Options{
			Path:     "/", // TODO check what this means
			MaxAge:   3600 * 24 * 30,
			HttpOnly: true,
		},
	}
	var err error
	a.Handler, err = auth.NewHandler(a,
		auth.SetStore(a.store),
		auth.SetSessionName(sessionName),
		auth.SetLanding(authLanding),
		auth.SetLogout(authLogout),
	)

	return err
}

// used by AuthenticateRequest
// TODO: should i trust GetCurrentUser() == nil for auth?
func (a *Auth) Check(name, pw string) (interface{}, error) {
	var u models.User
	err := db.DB.Where(&models.User{Email: name}).First(&u).Error
	if err != nil {
		log.Print("Login failed.", map[string]interface{}{
			"user": name,
			"err":  err,
		})
		return nil, auth.ErrBadLogin
	}
	err = bcrypt.CompareHashAndPassword(u.Hashed, []byte(pw))
	if err != nil {
		return nil, auth.ErrBadLogin
	}
	return u.ID, nil
}

func (Auth) LoginURL(c *admin.Context) string  { return "/login" }
func (Auth) LogoutURL(c *admin.Context) string { return authLogout }

func (a *Auth) GetCurrentUser(c *admin.Context) qor.CurrentUser {
	data, err := a.Handler.AuthenticateRequest(c.Context.Request)
	if err != nil {
		log.Print("qor/auth/GetCurrentUser: Authentication failed:", err)
		return nil
	}

	id, ok := data.(uint)
	if !ok {
		log.Print("qor/auth/GetCurrentUser: casting session failed.", map[string]interface{}{
			"data": data,
			"type": fmt.Sprintf("%t", data),
		})
		return nil
	}
	var u models.User
	u.ID = id
	err = db.DB.Where(u).First(&u).Error
	if err != nil {
		log.Print("qor/auth/GetCurrentUser: db lookup by ID failed", map[string]interface{}{
			"user": id,
			"err":  err,
		})
		return nil
	}
	return &u
}
