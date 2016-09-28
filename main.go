package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/csrf"
	"github.com/qor/qor-example/config"
	"github.com/qor/qor-example/config/admin"
	"github.com/qor/qor-example/config/api"
	_ "github.com/qor/qor-example/config/i18n"
	"github.com/qor/qor-example/config/routes"
	_ "github.com/qor/qor-example/db/migrations"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", routes.Router())
	admin.Admin.MountTo("/admin", mux)
	admin.Widgets.WidgetSettingResource.IndexAttrs("Name")

	api.API.MountTo("/api", mux)
	admin.Filebox.MountTo("/downloads", mux)

	for _, path := range []string{"system", "javascripts", "stylesheets", "images"} {
		mux.Handle(fmt.Sprintf("/%s/", path), http.FileServer(http.Dir("public")))
	}

	fmt.Printf("Listening on: %v\n", config.Config.Port)
	handler := csrf.Protect([]byte("3693f371bf91487c99286a777811bd4e"), csrf.Secure(false))(mux)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", config.Config.Port), handler); err != nil {
		panic(err)
	}
}
