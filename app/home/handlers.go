package home

import (
	"net/http"

	"github.com/qor/qor-example/config"
)

// Controller home controller
type Controller struct {
}

// Index home index page
func (Controller) Index(w http.ResponseWriter, req *http.Request) {
	config.View.Execute("index", map[string]interface{}{}, req, w)
}
