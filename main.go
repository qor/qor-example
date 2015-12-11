package main

import (
	"fmt"
	"net/http"

	"github.com/qor/qor-example/config"
	"github.com/qor/qor-example/config/admin"
	"github.com/qor/qor-example/config/routes"
	_ "github.com/qor/qor-example/db/migrations"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", routes.Rounter())
	admin.Admin.MountTo("/admin", mux)

	for _, path := range []string{"system", "downloads", "javascripts", "stylesheets", "images"} {
		mux.Handle(fmt.Sprintf("/%s/", path), http.FileServer(http.Dir("public")))
	}

	fmt.Printf("Listening on: %v\n", config.Config.Port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", config.Config.Port), mux); err != nil {
		panic(err)
	}
}
