package auth

import (
	"github.com/qor/auth"
	"github.com/qor/auth_themes/clean"
	"github.com/qor/qor-example/app/models"
	"github.com/qor/qor-example/config"
	"github.com/qor/qor-example/db"
)

// Auth initialize auth
var Auth = clean.New(&auth.Config{
	DB:        db.DB,
	Render:    config.View,
	UserModel: models.User{},
})
