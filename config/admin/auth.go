package admin

import (
	"github.com/qor/admin"
	"github.com/qor/qor"
	"github.com/qor/qor-example/app/models"
	"github.com/qor/qor-example/config/auth"
)

type Auth struct{}

func (Auth) LoginURL(c *admin.Context) string {
	return "/auth/login"
}

func (Auth) LogoutURL(c *admin.Context) string {
	return "/auth/logout"
}

func (Auth) GetCurrentUser(c *admin.Context) qor.CurrentUser {
	userInter, err := auth.Auth.CurrentUser(c.Writer, c.Request)
	if userInter != nil && err == nil {
		return userInter.(*models.User)
	}
	return nil
}
