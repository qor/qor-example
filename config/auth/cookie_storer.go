package auth

import (
	"encoding/base64"
	"fmt"
	"github.com/gorilla/securecookie"
	"gopkg.in/authboss.v0"
	"net/http"
	"time"
)

var CookieStore *securecookie.SecureCookie

func init() {
	cookieStoreKey, _ := base64.StdEncoding.DecodeString(`MLxWpWgQifDJ33zoCjz77rH5HR3ifJ8kmT3JqfUu_CEKbGXkwCXK_RoyKCnwoEd1-WnRP2cFcDZxxycpQLp3EQ==`)
	CookieStore = securecookie.New(cookieStoreKey, nil)
}

type CookieStorer struct {
	w http.ResponseWriter
	r *http.Request
}

func NewCookieStorer(w http.ResponseWriter, r *http.Request) authboss.ClientStorer {
	return &CookieStorer{w, r}
}

func (s CookieStorer) Get(key string) (string, bool) {
	cookie, err := s.r.Cookie(key)
	if err != nil {
		return "", false
	}

	var value string
	err = CookieStore.Decode(key, cookie.Value, &value)
	if err != nil {
		return "", false
	}

	return value, true
}

func (s CookieStorer) Put(key, value string) {
	encoded, err := CookieStore.Encode(key, value)
	if err != nil {
		fmt.Println(err)
	}

	cookie := &http.Cookie{
		Expires: time.Now().UTC().AddDate(1, 0, 0),
		Name:    key,
		Value:   encoded,
		Path:    "/",
	}
	http.SetCookie(s.w, cookie)
}

func (s CookieStorer) Del(key string) {
	cookie := &http.Cookie{
		MaxAge: -1,
		Name:   key,
		Path:   "/",
	}
	http.SetCookie(s.w, cookie)
}
