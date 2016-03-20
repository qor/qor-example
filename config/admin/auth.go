package admin

import (
	"github.com/gorilla/sessions"
	"github.com/qor/admin"
	"github.com/qor/qor"
	"github.com/qor/qor-example/app/models"
	"github.com/qor/qor-example/config"
	"github.com/qor/qor-example/db"
)

type Auth struct {
	User     string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func (Auth) LoginURL(c *admin.Context) string {
	return "/login"
}

func (Auth) LogoutURL(c *admin.Context) string {
	return "/login"
}

func (Auth) GetCurrentUser(c *admin.Context) qor.CurrentUser {
	var store = sessions.NewCookieStore([]byte(config.Config.Secret))
	session, err := store.Get(c.Request, config.Config.Session.Name)
	if err != nil {
		return nil
	}
	var currentUser models.User
	if session.Values["_auth_user_id"] != nil {
		if !c.GetDB().Where("id = ?", session.Values["_auth_user_id"]).First(&currentUser).RecordNotFound() {
			return &currentUser
		}
	}
	return nil

	// OR
	// return &models.User{Name: "Admin"}
}

// Return User
func (this *Auth) GetUser() (bool, *models.User) {
	var currentUser models.User
	if !db.DB.Where("name = ? OR email = ?", this.User, this.User).First(&currentUser).RecordNotFound() {
		return true, &currentUser
	}
	return false, &currentUser
}
