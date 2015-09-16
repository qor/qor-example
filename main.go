package main

import (
	"fmt"
	"net/http"

	"github.com/grengojbo/qor-example/config"
	"github.com/grengojbo/qor-example/config/admin"
	_ "github.com/grengojbo/qor-example/db/migrations"
)

func main() {
	mux := http.NewServeMux()
	admin.Admin.MountTo("/admin", mux)

	mux.Handle("/system/", http.FileServer(http.Dir("public")))

	fmt.Printf("Listening on: %v\n", config.Config.Port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", config.Config.Port), mux); err != nil {
		panic(err)
	}
}
