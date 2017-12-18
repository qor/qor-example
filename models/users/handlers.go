package users

import "net/http"

// Controller User Controller
type Controller struct {
}

func (Controller) Profile(w http.ResponseWriter, req *http.Request) {
}

func (Controller) Orders(w http.ResponseWriter, req *http.Request) {
}

func (Controller) Update(w http.ResponseWriter, req *http.Request) {
}

func (Controller) AddCredit(w http.ResponseWriter, req *http.Request) {
}
