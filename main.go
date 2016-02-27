package main

import (
	"fmt"
	"net/http"

	//go:generate go-bindata -nomemcopy ../qor/admin/views/...
	// "github.com/gin-gonic/contrib/sessions"
	// "github.com/grengojbo/gotools"
	"github.com/gin-gonic/gin"
	"github.com/qor/qor-example/config"
	"github.com/qor/qor-example/config/admin"
	"github.com/qor/qor-example/config/routes"
	_ "github.com/qor/qor-example/db/migrations"
)

var (
	Version   = "0.1.0"
	BuildTime = "2015-09-20 UTC"
	GitHash   = "c00"
)

func main() {
	conf := config.Config
	fmt.Printf("App Version: %s\n", Version)
	fmt.Printf("Build Time: %s\n", BuildTime)
	fmt.Printf("Git Commit Hash: %s\n", GitHash)
	fmt.Printf("Listening on: %v\n", conf.Port)

	mux := http.NewServeMux()
	// mux.Handle("/", routes.Router())
	admin.Admin.MountTo("/admin", mux)
	// api.API.MountTo("/api", mux)

	r := routes.Router()
	r.Any("/admin/*w", gin.WrapH(mux))
	r.Run(fmt.Sprintf(":%d", conf.Port))
}
