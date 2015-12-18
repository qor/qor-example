package api

import (
	"github.com/qor/qor"
	"github.com/qor/qor-example/app/models"
	"github.com/qor/qor-example/db"
	"github.com/qor/qor/admin"
)

var API *admin.Admin

func init() {
	API = admin.New(&qor.Config{DB: db.DB})

	API.AddResource(&models.Product{})
	API.AddResource(&models.Order{})
	API.AddResource(&models.User{})
}
