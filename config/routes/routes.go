package routes

import (
	"context"
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi"
	"github.com/qor/publish2"
	"github.com/qor/qor"
	"github.com/qor/qor/utils"
	"github.com/qor/wildcard_router"

	"github.com/qor/qor-example/app/controllers"
	"github.com/qor/qor-example/config"
	"github.com/qor/qor-example/config/admin/bindatafs"
	"github.com/qor/qor-example/config/auth"
	"github.com/qor/qor-example/db"
)

var rootMux *http.ServeMux
var WildcardRouter *wildcard_router.WildcardRouter

func Router() *http.ServeMux {
	if rootMux == nil {
		router := chi.NewRouter()

		router.Use(func(next http.Handler) http.Handler {
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

		router.Get("/", controllers.HomeIndex)
		router.Get("/products", controllers.ProductIndex)
		router.Get("/products/{code}", controllers.ProductShow)
		router.Get("/g/{gender}", controllers.ProductGenderShow)
		router.Get("/category/{code}", controllers.CategoryShow)
		router.Get("/switch_locale", controllers.SwitchLocale)

		router.With(auth.Authority.Authorize()).Route("/account", func(r chi.Router) {
			r.Get("/", controllers.AccountShow)
			r.With(auth.Authority.Authorize("logged_in_half_hour")).Post("/add_user_credit", controllers.AddUserCredit)
			r.Get("/profile", controllers.ProfileShow)
			r.Post("/profile", controllers.SetUserProfile)
			r.Post("/profile/billing_address", controllers.SetBillingAddress)
			r.Post("/profile/shipping_address", controllers.SetShippingAddress)
		})

		router.Route("/cart", func(r chi.Router) {
			r.Get("/", controllers.ShowCartHandler)
			r.Get("/checkout", controllers.CheckoutCartHandler)
			// r.Post("/", controllers.AddToCartHandler)
			// r.Post("/checkout", controllers.OrderCartHandler)
			// r.Delete("/:id", controllers.RemoveFromCartHandler)
		})

		rootMux = http.NewServeMux()

		rootMux.Handle("/auth/", auth.Auth.NewServeMux())
		rootMux.Handle("/system/", utils.FileServer(http.Dir(filepath.Join(config.Root, "public"))))
		assetFS := bindatafs.AssetFS.FileServer(http.Dir("public"), "javascripts", "stylesheets", "images", "dist", "fonts", "vendors")
		for _, path := range []string{"javascripts", "stylesheets", "images", "dist", "fonts", "vendors"} {
			rootMux.Handle(fmt.Sprintf("/%s/", path), assetFS)
		}

		WildcardRouter = wildcard_router.New()
		WildcardRouter.MountTo("/", rootMux)
		WildcardRouter.AddHandler(router)
	}
	return rootMux
}
