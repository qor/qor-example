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
	"github.com/qor/qor-example/config/api"
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
	api.API.MountTo("/api", mux)

	// r := gin.Default()
	// if conf.Session.Adapter == "redis" {
	// 	store, err := sessions.NewRedisStore(10, conf.Redis.Protocol, fmt.Sprintf("%v:%v", conf.Redis.Host, conf.Redis.Port), "", []byte(conf.Secret))
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	r.Use(sessions.Sessions(conf.Session.Name, store))
	// } else if conf.Session.Adapter == "cookie" {
	// 	store := sessions.NewCookieStore([]byte(conf.Secret))
	// 	r.Use(sessions.Sessions(conf.Session.Name, store))
	// }
	// r.LoadHTMLGlob("app/views/*.tmpl")
	// // r.LoadHTMLGlob("admin/views/*.tmpl")

	// for _, path := range []string{"static", "system", "downloads", "javascripts", "stylesheets", "images"} {
	// 	r.Static(fmt.Sprintf("/%s", path), fmt.Sprintf("public/%s", path))
	// }

	// r.GET("/logout", func(c *gin.Context) {
	// 	session := sessions.Default(c)
	// 	session.Clear()
	// 	session.Save()
	// 	c.Redirect(http.StatusMovedPermanently, "/login")
	// })

	// r.GET("/login", func(c *gin.Context) {
	// 	session := sessions.Default(c)
	// 	session.Set("lastLogin", time.Now().Unix())
	// 	session.Set("ip", c.ClientIP())
	// 	session.Save()
	// 	c.HTML(200, "login.tmpl", gin.H{
	// 		"title":     admin.Admin.SiteName,
	// 		"timestamp": time.Now().Unix(),
	// 	})
	// })

	// r.POST("/login", func(c *gin.Context) {
	// 	var login admin.Auth
	// 	session := sessions.Default(c)
	// 	if c.BindJSON(&login) == nil {
	// 		if ok, user := login.GetUser(); ok != false {
	// 			if err := gotools.VerifyPassword(user.Password, login.Password); err != nil {
	// 				session.Set("lastLogin", time.Now().Unix())
	// 				session.Save()
	// 				c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized", "message": "User unauthorized"})
	// 			} else {
	// 				session.Set("lastLogin", time.Now().Unix())
	// 				session.Set("_auth_user_id", user.ID)
	// 				session.Save()
	// 				c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Ok"})
	// 			}
	// 		} else {
	// 			session.Set("lastLogin", time.Now().Unix())
	// 			session.Save()
	// 			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized", "message": "User unauthorized"})
	// 		}
	// 	} else {
	// 		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Bad request"})
	// 	}
	// })

	// r.GET("/", func(c *gin.Context) {
	// 	c.Redirect(http.StatusMovedPermanently, "/admin")
	// })

	r := routes.Router()
	r.Any("/admin/*w", gin.WrapH(mux))
	r.Any("/api/*w", gin.WrapH(mux))
	r.Run(fmt.Sprintf(":%d", conf.Port))
}
