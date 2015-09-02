package main

import (
	"fmt"
	"net/http"

	"github.com/qor/qor-example/config"
	"github.com/qor/qor-example/config/admin"
	_ "github.com/qor/qor-example/db/migrations"
)

func main() {
	mux := http.NewServeMux()
	admin.Admin.MountTo("/admin", mux)

	fmt.Printf("Listening on: %v\n", config.Config.Port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", config.Config.Port), mux); err != nil {
		panic(err)
	}
}
