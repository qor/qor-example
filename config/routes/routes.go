package routes

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/grengojbo/gotools"
	"github.com/itsjamie/gin-cors"
	"github.com/qor/qor-example/app/controllers"
	"github.com/qor/qor-example/config"
	"github.com/qor/qor-example/config/admin"
)

func Router() *gin.Engine {
	conf := config.Config
	router := gin.Default()
	gin.SetMode(gin.DebugMode)
	router.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))

	if conf.Session.Adapter == "redis" {
		store, err := sessions.NewRedisStore(10, conf.Redis.Protocol, fmt.Sprintf("%v:%v", conf.Redis.Host, conf.Redis.Port), "", []byte(conf.Secret))
		if err != nil {
			panic(err)
		}
		router.Use(sessions.Sessions(conf.Session.Name, store))
	} else if conf.Session.Adapter == "cookie" {
		store := sessions.NewCookieStore([]byte(conf.Secret))
		router.Use(sessions.Sessions(conf.Session.Name, store))
	}

	// for _, path := range []string{"static", "downloads"} {
	for _, path := range []string{"static", "system", "downloads", "javascripts", "stylesheets", "images", "css", "fonts", "js"} {
		router.Static(fmt.Sprintf("/%s", path), fmt.Sprintf("public/%s", path))
	}

	// r.LoadHTMLGlob("app/views/*.tmpl")
	if tmpl, err := template.New("projectViews").Funcs(config.FuncMap).ParseGlob("app/views/*.tmpl"); err == nil {
		router.SetHTMLTemplate(tmpl)
	} else {
		panic(err)
	}
	router.GET("/", controllers.HomeIndex)
	router.GET("/products", controllers.ProductIndex)
	router.GET("/products/:code", controllers.ProductShow)
	// router.HandleFunc("/guitars/{id:[0-9]+}", h.guitarsShowHandler).Methods("GET")

	// API version 1
	v1 := router.Group("api/v1")
	v1.GET("/category", controllers.CategoryIndex)
	v1.GET("/products", controllers.ProductApiIndex)
	v1.GET("/orders", controllers.OrderIndex)
	v1.GET("/users/:id", controllers.UserShow)
	v1.POST("/auth/login", controllers.LoginApi)
	v1.POST("/auth/logout", controllers.LogoutApi)

	// router.GET("/", func(c *gin.Context) {
	// 	c.Redirect(http.StatusMovedPermanently, "/admin")
	// })

	router.GET("/logout", func(c *gin.Context) {
		session := sessions.Default(c)
		session.Clear()
		session.Save()
		c.Redirect(http.StatusMovedPermanently, "/login")
	})

	router.GET("/login", func(c *gin.Context) {
		session := sessions.Default(c)
		session.Set("lastLogin", time.Now().Unix())
		session.Set("ip", c.ClientIP())
		session.Save()
		c.HTML(200, "login.tmpl", gin.H{
			"title":     conf.SiteName,
			"timestamp": time.Now().Unix(),
		})
	})

	router.POST("/login", func(c *gin.Context) {
		var login admin.Auth
		session := sessions.Default(c)
		if c.BindJSON(&login) == nil {
			if ok, user := login.GetUser(); ok != false {
				if err := gotools.VerifyPassword(user.Password, login.Password); err != nil {
					session.Set("lastLogin", time.Now().Unix())
					session.Save()
					c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized", "message": "User unauthorized"})
				} else {
					session.Set("lastLogin", time.Now().Unix())
					session.Set("_auth_user_id", user.ID)
					session.Save()
					c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Ok"})
				}
			} else {
				session.Set("lastLogin", time.Now().Unix())
				session.Save()
				c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized", "message": "User unauthorized"})
			}
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Bad request"})
		}
	})
	return router
}
