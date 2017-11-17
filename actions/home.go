package actions

import (
	"github.com/gobuffalo/buffalo"
)

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {

	s := c.Session()
	c.Set("uid", s.Get("uid"))
	c.Set("username", s.Get("username"))
	c.Set("avatar", s.Get("avatar"))

	m := make([]map[string]string, 0)
	m = append(m, map[string]string{"gplus": "google"})
	c.Set("Providers", m)
	return c.Render(200, r.HTML("index.html"))
}
