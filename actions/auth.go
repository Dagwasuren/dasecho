package actions

import (
	"github.com/gobuffalo/buffalo"
)

func AuthFacebookHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("index.html"))
}
func AuthGithubHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("index.html"))
}
func AuthTwitterHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("index.html"))
}
func AuthGoogleHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("index.html"))
}
