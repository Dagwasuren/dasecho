package actions

import (
	"log"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/mw-paramlogger"
	"github.com/gobuffalo/buffalo/mw-forcessl"
	"github.com/gobuffalo/envy"
	"github.com/unrolled/secure"

	"github.com/netroby/dasecho/models"
	"github.com/gobuffalo/buffalo/mw-csrf"
	"github.com/gobuffalo/buffalo/mw-i18n"
	"github.com/gobuffalo/packr"
	"github.com/gobuffalo/buffalo-pop/pop/popmw"

	"github.com/gorilla/sessions"
	"github.com/markbates/goth/gothic"
)

// ENV is used to help switch settings based on where the
// application is being run. Default is "development".
var ENV = envy.Get("GO_ENV", "development")
var app *buffalo.App
var T *i18n.Translator

// App is where all routes and middleware for buffalo
// should be defined. This is the nerve center of your
// application.
func App() *buffalo.App {
	appHost := "http://dev.dasecho.net:3000"
	if ENV == "production" {
		appHost = "https://dasecho.net"
	}
	if app == nil {
		app = buffalo.New(buffalo.Options{
			Env:          ENV,
			Host:         appHost,
			SessionName:  "_dasecho_session",
			SessionStore: sessions.NewCookieStore([]byte(envy.Get("SSKEY", "DKSLFAJ7234DSDFJSAOOZNWEROZ"))),
		})
		// Automatically redirect to SSL
		app.Use(ssl.ForceSSL(secure.Options{
			SSLRedirect:     ENV == "production",
			SSLProxyHeaders: map[string]string{"X-Forwarded-Proto": "https"},
		}))

		// Automatically save the session if the underlying
		// Handler does not return an error.

		if ENV == "development" {
			app.Use(paramlogger.ParameterLogger)
		}

		if ENV != "test" {
			// Protect against CSRF attacks. https://www.owasp.org/index.php/Cross-Site_Request_Forgery_(CSRF)
			// Remove to disable this.
			app.Use(csrf.New)
		}

		// Wraps each request in a transaction.
		//  c.Value("tx").(*pop.PopTransaction)
		// Remove to disable this.
		app.Use(popmw.Transaction(models.DB))

		// Setup and use translations:
		var err error
		if T, err = i18n.New(packr.NewBox("../locales"), "en-US"); err != nil {
			app.Stop(err)
		}
		app.Use(T.Middleware())

		app.GET("/", HomeHandler)

		app.ServeFiles("/assets", packr.NewBox("../public/assets"))
		log.Println("Initialized Auth")
		log.Println("Get webhost ", app.Host)
		auth := app.Group("/auth")
		auth.GET("/{provider}", buffalo.WrapHandlerFunc(gothic.BeginAuthHandler))
		auth.GET("/{provider}/callback", AuthCallback)
		app.GET("/article/create", ArticleCreate)
		app.POST("/article/save-create", ArticleSaveCreate)
		app.GET("/article/edit", ArticleEdit)
		app.POST("/article/save-edit", ArticleSaveEdit)
		app.GET("/article/delete", ArticleDelete)
		app.GET("/todaybest/create", TodaybestCreate)
		app.POST("/todaybest/save-create", TodaybestSaveCreate)
		app.GET("/todaybest/edit", TodaybestEdit)
		app.GET("/todaybest/list", TodaybestList)
		app.POST("/todaybest/save-edit", TodaybestSaveEdit)
		app.GET("/todaybest/delete", TodaybestDelete)
	}

	return app
}
