package actions

import (
	"github.com/gobuffalo/buffalo"
)

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {

	m := make([]map[string]string, 0)

	m = append(m, map[string]string{"facebook": "facebook"})
	m = append(m, map[string]string{"github": "github"})
	m = append(m, map[string]string{"twitter": "twitter"})
	m = append(m, map[string]string{"google": "google"})
	c.Set("Providers", m)
	return c.Render(200, r.HTML("index.html"))
}
