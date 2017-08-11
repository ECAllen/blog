package actions

import (
	"fmt"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/middleware"
	"github.com/gobuffalo/buffalo/middleware/csrf"
	"github.com/gobuffalo/buffalo/middleware/i18n"

	"github.com/gobuffalo/envy"
	"github.com/gobuffalo/packr"
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
	if app == nil {
		app = buffalo.Automatic(buffalo.Options{
			Env:         ENV,
			SessionName: "_blog_session",
			Host:        "http://ecallen.com",
		})

		if ENV == "development" {
			app.Use(middleware.ParameterLogger)
		}
		if ENV != "test" {
			// Protect against CSRF attacks. https://www.owasp.org/index.php/Cross-Site_Request_Forgery_(CSRF)
			// Remove to disable this.
			app.Use(csrf.Middleware)
		}

		// Setup and use translations:
		var err error
		if T, err = i18n.New(packr.NewBox("../locales"), "en-US"); err != nil {
			app.Stop(err)
		}
		app.Use(T.Middleware())

		app.Use(SetVars)

		app.GET("/", HomeHandler)

		app.GET("/blog-posts/posts/{post}", func(c buffalo.Context) error {
			return c.Render(200, r.HTML("blog-posts/posts/"+c.Param("post")))
		})

		app.GET("/notes/index.md", func(c buffalo.Context) error {
			return c.Render(200, r.HTML("/notes/index.md"))
		})

		notes := app.Group("/notes")
		notes.GET("/{section}/{post}", func(c buffalo.Context) error {
			path := fmt.Sprintf("%v", c.Value("current_path"))
			return c.Render(200, r.HTML(path))
		})

		/*
			app.GET("/notes/{section}/{post}", func(c buffalo.Context) error {
				return c.Render(200, r.HTML("/notes/"+c.Param("section")+"/"+c.Param("section")))
			})
		*/

		app.ServeFiles("/assets", packr.NewBox("../public/assets"))
	}

	return app
}
