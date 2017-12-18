package home

import (
	"net/http"

	"github.com/qor/qor-example/config"
)

// Index home index page
func Index(w http.ResponseWriter, req *http.Request) {
	config.View.Execute("home_index", map[string]interface{}{}, req, w)
}
