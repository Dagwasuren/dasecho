package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/markbates/pop"
	"github.com/dasecho/dasecho/models"
	"github.com/pkg/errors"
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

	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)
	articles := &models.Articles{}
	// You can order your list here. Just change
	err := tx.Order("id desc").All(articles)
	// to:
	// err := tx.Order("create_at desc").All(articles)
	if err != nil {
		return errors.WithStack(err)
	}
	// Make articles available inside the html template
	c.Set("articles", articles)
	
	return c.Render(200, r.HTML("index.html"))
}
