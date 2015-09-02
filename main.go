package main

import (
	"fmt"
	"net/http"

	"github.com/qor/qor-example/config"
)

func main() {
	mux := http.NewServeMux()
	// Admin.MountTo("/admin", mux)

	// start the server
	fmt.Printf("Listening on: %v\n", config.Config.Port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", config.Config.Port), mux); err != nil {
		panic(err)
	}
}
