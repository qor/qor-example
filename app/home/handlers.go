package home

import (
	"net/http"

	"github.com/qor/qor-example/config"
)

// Handler home handlers
type Handler struct {
}

// Index home index page
func (Handler) Index(w http.ResponseWriter, req *http.Request) {
	config.View.Execute("index", map[string]interface{}{}, req, w)
}
