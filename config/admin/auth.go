package admin

import (
	"github.com/grengojbo/qor-example/app/models"
	"github.com/qor/qor"
	"github.com/qor/qor/admin"
)

type Auth struct{}

func (Auth) LoginURL(c *admin.Context) string {
	return "/admin"
}

func (Auth) LogoutURL(c *admin.Context) string {
	return "/admin"
}

func (Auth) GetCurrentUser(c *admin.Context) qor.CurrentUser {
	return &models.User{Name: "Admin"}
}
