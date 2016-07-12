package auth

import (
	"encoding/base64"
	"fmt"
	"github.com/gorilla/sessions"
	"gopkg.in/authboss.v0"
	"net/http"
)

const sessionCookieName = "qor-example"

var sessionStore *sessions.CookieStore

type SessionStorer struct {
	w http.ResponseWriter
	r *http.Request
}

func init() {
	sessionStoreKey, _ := base64.StdEncoding.DecodeString(`EGxRGrqkogymr5EHAX9V3RAzaZzj9_heLXM4M6DZDuGsEd-nfw8veekXWDY11pfsWXtQsJMzZxRm2zpjGg9dJQ==`)
	sessionStore = sessions.NewCookieStore(sessionStoreKey)
}

func NewSessionStorer(w http.ResponseWriter, r *http.Request) authboss.ClientStorer {
	return &SessionStorer{w, r}
}

func (s SessionStorer) Get(key string) (string, bool) {
	session, err := sessionStore.Get(s.r, sessionCookieName)
	if err != nil {
		fmt.Println(err)
		return "", false
	}

	strInf, ok := session.Values[key]
	if !ok {
		return "", false
	}

	str, ok := strInf.(string)
	if !ok {
		return "", false
	}

	return str, true
}

func (s SessionStorer) Put(key, value string) {
	session, err := sessionStore.Get(s.r, sessionCookieName)
	if err != nil {
		fmt.Println(err)
		return
	}

	session.Values[key] = value
	session.Save(s.r, s.w)
}

func (s SessionStorer) Del(key string) {
	session, err := sessionStore.Get(s.r, sessionCookieName)
	if err != nil {
		fmt.Println(err)
		return
	}

	delete(session.Values, key)
	session.Save(s.r, s.w)
}
