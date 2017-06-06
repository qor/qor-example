package main

import (
	"flag"
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"github.com/gorilla/csrf"
	"github.com/qor/action_bar"
	"github.com/qor/i18n/inline_edit"
	"github.com/qor/qor"
	"github.com/qor/qor-example/app/models"
	"github.com/qor/qor-example/config"
	"github.com/qor/qor-example/config/admin"
	"github.com/qor/qor-example/config/admin/bindatafs"
	"github.com/qor/qor-example/config/api"
	"github.com/qor/qor-example/config/i18n"
	"github.com/qor/qor-example/config/routes"
	"github.com/qor/qor-example/config/seo"
	"github.com/qor/qor-example/config/utils"
	"github.com/qor/qor-example/db"
	_ "github.com/qor/qor-example/db/migrations"
	"github.com/qor/render"
)

func main() {
	var compileTemplate = flag.Bool("compile-templates", false, "Compile Templates")
	flag.Parse()

	mux := http.NewServeMux()
	mux.Handle("/", routes.Router())
	admin.Admin.MountTo("/admin", mux)
	admin.Filebox.MountTo("/downloads", mux)
	api.API.MountTo("/api", mux)

	config.View.FuncMapMaker = func(render *render.Render, request *http.Request, writer http.ResponseWriter) template.FuncMap {
		funcMap := template.FuncMap{}

		// Add `t` method
		for key, fc := range inline_edit.FuncMap(i18n.I18n, utils.GetCurrentLocale(request), utils.GetEditMode(writer, request)) {
			funcMap[key] = fc
		}

		// Add `action_bar` method
		funcMap["render_action_bar"] = func() template.HTML {
			return admin.ActionBar.Actions(action_bar.Action{Name: "Edit SEO", Link: seo.SEOCollection.SEOSettingURL("/help")}).Render(writer, request)
		}

		funcMap["render_seo_tag"] = func() template.HTML {
			// FIXME get db
			return seo.SEOCollection.Render(&qor.Context{DB: db.DB}, "Default Page")
		}

		funcMap["get_categories"] = func() (categories []models.Category) {
			// FIXME get db
			db.DB.Find(&categories)
			return
		}

		funcMap["current_locale"] = func() string {
			return utils.GetCurrentLocale(request)
		}

		funcMap["current_user"] = func() *models.User {
			return nil
		}

		return funcMap
	}

	skipCheck := func(h http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			if !strings.HasPrefix(r.URL.Path, "/auth") {
				r = csrf.UnsafeSkipCheck(r)
			}
			h.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
	handler := csrf.Protect([]byte("3693f371bf91487c99286a777811bd4e"), csrf.Secure(false))(mux)

	if *compileTemplate {
		bindatafs.AssetFS.Compile()
	} else {
		fmt.Printf("Listening on: %v\n", config.Config.Port)
		if err := http.ListenAndServe(fmt.Sprintf(":%d", config.Config.Port), skipCheck(handler)); err != nil {
			panic(err)
		}
	}
}
