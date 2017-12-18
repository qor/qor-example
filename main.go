package main

import (
	"github.com/go-chi/chi"
	"github.com/qor/admin"
	"github.com/qor/middlewares"
	"github.com/qor/publish2"
	"github.com/qor/qor-example/app/home"
	"github.com/qor/qor-example/config/admin/bindatafs"
	"github.com/qor/qor-example/config/application"
	"github.com/qor/qor-example/config/db"

	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/qor/qor"
	"github.com/qor/qor-example/config"
	_ "github.com/qor/qor-example/config/db/migrations"
)

func main() {
	cmdLine := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	compileTemplate := cmdLine.Bool("compile-templates", false, "Compile Templates")
	cmdLine.Parse(os.Args[1:])

	var (
		Router      = chi.NewRouter()
		Admin       = admin.New(&qor.Config{DB: db.DB.Set(publish2.VisibleMode, publish2.ModeOff).Set(publish2.ScheduleMode, publish2.ModeOff)})
		Application = application.New(&application.Config{
			Router: Router,
			Admin:  Admin,
			DB:     db.DB,
		})
	)

	Application.Use(home.New(&home.Config{}))
	if *compileTemplate {
		bindatafs.AssetFS.Compile()
	} else {
		fmt.Printf("Listening on: %v\n", config.Config.Port)
		if err := http.ListenAndServe(fmt.Sprintf(":%d", config.Config.Port), middlewares.Apply(Router)); err != nil {
			panic(err)
		}
	}
}
