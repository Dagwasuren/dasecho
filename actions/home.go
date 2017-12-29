package actions

import (
	"log"

	"github.com/dasecho/dasecho/models"
	"github.com/gobuffalo/buffalo"
	"github.com/markbates/pop"
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

	// Paginate results. Params "page" and "per_page" control     pagination.
	// Default values are "page=1" and "per_page=2".
	q := tx.PaginateFromParams(c.Params())
	err := q.Order("created_at desc").All(articles)
	if err != nil {
		return errors.WithStack(err)
	}
	// Add the paginator to the context so it can be used in the template.
	c.Set("pagination", q.Paginator)

	// Make articles available inside the html template
	c.Set("articles", articles)
	todaybest := &models.Todaybest{}
	err = tx.Order("created_at desc").First(todaybest)
	if err != nil {
		log.Println(err)
		c.Set("todaybest", nil)
	} else {
		c.Set("todaybest", todaybest)
	}

	return c.Render(200, r.HTML("index.html"))
}
