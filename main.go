package main

import "net/http"

func main() {
	mux := http.NewServeMux()
	// Admin.MountTo("/admin", mux)

	// start the server
	http.ListenAndServe(":9000", mux)
}
