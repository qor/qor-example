package main

import (
	"fmt"
	"net/http"

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

	// subrouter - protected by admin.QorAuth.Handler.Authenticate
	adminMux := http.NewServeMux()

	admin.Admin.MountTo("/admin/qor", adminMux)

	// todo: move this to frontend templating
	mux.Handle("/admin_login", http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			admin.QorAuth.Handler.Authorize(rw, r)
		} else {
			fmt.Fprintf(rw, `
<html>
    <body>
    <h1>Login</h1>
    <form method="POST">
        <input type="text" name="user">
        <input type="password" name="pass">
        <input type="submit">
    </form>
    </body>
</html>`)
		}
	}))
	mux.Handle("/admin_logout", http.HandlerFunc(admin.QorAuth.Handler.Logout))

	mux.Handle("/admin/", admin.QorAuth.Handler.Authenticate(adminMux))
	api.API.MountTo("/api", mux)

	for _, path := range []string{"system", "downloads", "javascripts", "stylesheets", "images"} {
		mux.Handle(fmt.Sprintf("/%s/", path), http.FileServer(http.Dir("public")))
	}

	fmt.Printf("Listening on: %v\n", config.Config.Port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", config.Config.Port), mux); err != nil {
		panic(err)
	}
}
