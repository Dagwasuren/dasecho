package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/markbates/pop"
	"github.com/dasecho/dasecho/models"
	"github.com/pkg/errors"
)

// ArticleCreate default implementation.
func ArticleCreate(c buffalo.Context) error {
	s := c.Session()
	username := s.Get("username")
	if username == nil {
		c.Set("message", "请先登录")
		return c.Render(422, r.HTML("message.html"))
	}
	return c.Render(200, r.HTML("article/create.html"))
}

func ArticleSaveCreate(c buffalo.Context) error {
	a := &models.Article{}
	c.Request().ParseForm()
	a.Title = c.Request().Form.Get("title")
	a.Content = c.Request().Form.Get("content")
	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)
	s := c.Session()
	username := s.Get("username")
	if username != nil {
		a.Author = username.(string)
	} else {
		c.Set("message", "请先登录")
		return c.Render(422, r.HTML("message.html"))
	}
	uid := s.Get("uid")
	if uid != nil {
		a.Uid = s.Get("uid").(int)
	} else {

		a.Uid = 1
	}
	// Validate the data from the html form
	verrs, err := tx.ValidateAndCreate(a)
	if err != nil {
		return errors.WithStack(err)
	}
	if verrs.HasAny() {
		c.Set("errors", verrs)
		return c.Render(422, r.HTML("message.html"))
	}
		c.Set("message", "Insert Success")
		// Render again the new.html template that the user can
		// correct the input.
		return c.Render(200, r.HTML("message.html"))
}
// ArticleEdit default implementation.
func ArticleEdit(c buffalo.Context) error {
	return c.Render(200, r.HTML("article/edit.html"))
}
func ArticleSaveEdit(c buffalo.Context) error {
	return c.Render(200, r.JSON("message"))
}

// ArticleDelete default implementation.
func ArticleDelete(c buffalo.Context) error {
	return c.Render(200, r.HTML("article/delete.html"))
}
