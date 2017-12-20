package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/qor/middlewares"
	"github.com/qor/publish2"
	"github.com/qor/qor"
	"github.com/qor/qor-example/app/account"
	"github.com/qor/qor-example/app/home"
	"github.com/qor/qor-example/app/orders"
	"github.com/qor/qor-example/app/pages"
	"github.com/qor/qor-example/app/products"
	"github.com/qor/qor-example/config"
	"github.com/qor/qor-example/config/admin"
	"github.com/qor/qor-example/config/admin/bindatafs"
	"github.com/qor/qor-example/config/api"
	"github.com/qor/qor-example/config/application"
	"github.com/qor/qor-example/config/db"
	_ "github.com/qor/qor-example/config/db/migrations"
	"github.com/qor/qor/utils"
	"github.com/qor/wildcard_router"
)

func main() {
	cmdLine := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	compileTemplate := cmdLine.Bool("compile-templates", false, "Compile Templates")
	cmdLine.Parse(os.Args[1:])

	var (
		Router      = chi.NewRouter()
		Admin       = admin.Admin // admin.New(&qor.Config{DB: db.DB.Set(publish2.VisibleMode, publish2.ModeOff).Set(publish2.ScheduleMode, publish2.ModeOff)})
		Application = application.New(&application.Config{
			Router: Router,
			Admin:  Admin,
			DB:     db.DB,
		})
	)

	Router.Use(middleware.RealIP)
	Router.Use(middleware.Logger)
	Router.Use(middleware.Recoverer)
	Router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			var (
				tx         = db.DB
				qorContext = &qor.Context{Request: req, Writer: w}
			)

			if locale := utils.GetLocale(qorContext); locale != "" {
				tx = tx.Set("l10n:locale", locale)
			}

			ctx := context.WithValue(req.Context(), utils.ContextDBName, publish2.PreviewByDB(tx, qorContext))
			next.ServeHTTP(w, req.WithContext(ctx))
		})
	})

	Application.Use(home.New(&home.Config{}))
	Application.Use(products.New(&products.Config{}))
	Application.Use(account.New(&account.Config{}))
	Application.Use(orders.New(&orders.Config{}))
	Application.Use(pages.New(&pages.Config{}))

	mux := http.NewServeMux()
	admin.Admin.MountTo("/admin", mux)
	api.API.MountTo("/api", mux)
	mux.Handle("/system/", utils.FileServer(http.Dir(filepath.Join(config.Root, "public"))))
	assetFS := bindatafs.AssetFS.FileServer(http.Dir("public"), "javascripts", "stylesheets", "images", "dist", "fonts", "vendors")
	for _, path := range []string{"javascripts", "stylesheets", "images", "dist", "fonts", "vendors"} {
		mux.Handle(fmt.Sprintf("/%s/", path), assetFS)
	}

	wildcardRouter := wildcard_router.New()
	wildcardRouter.AddHandler(Router)
	wildcardRouter.AddHandler(admin.MicroSite)
	wildcardRouter.MountTo("/", mux)

	if *compileTemplate {
		bindatafs.AssetFS.Compile()
	} else {
		fmt.Printf("Listening on: %v\n", config.Config.Port)
		if err := http.ListenAndServe(fmt.Sprintf(":%d", config.Config.Port), middlewares.Apply(mux)); err != nil {
			panic(err)
		}
	}
}
