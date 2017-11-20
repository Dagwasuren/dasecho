package actions

import (
	"strconv"

	"github.com/dasecho/dasecho/models"
	"github.com/gobuffalo/buffalo"
	"github.com/markbates/pop"
	"github.com/pkg/errors"
)

// ArticleCreate default implementation.
func ArticleCreate(c buffalo.Context) error {

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
	err := tx.Order("created_at desc").All(articles)
	// to:
	// err := tx.Order("create_at desc").All(articles)
	if err != nil {
		return errors.WithStack(err)
	}
	// Make articles available inside the html template
	c.Set("articles", articles)

	s = c.Session()
	username := s.Get("username")
	if username == nil {
		c.Set("message", "请先登录")
		return c.Render(422, r.HTML("message.html"))
	}
	return c.Render(200, r.HTML("article/create.html"))
}

func ArticleSaveCreate(c buffalo.Context) error {

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
	err := tx.Order("created_at desc").All(articles)
	// to:
	// err := tx.Order("create_at desc").All(articles)
	if err != nil {
		return errors.WithStack(err)
	}
	// Make articles available inside the html template
	c.Set("articles", articles)

	a := &models.Article{}
	c.Request().ParseForm()
	a.Title = c.Request().Form.Get("title")
	a.Content = c.Request().Form.Get("content")
	// Get the DB connection from the context
	tx = c.Value("tx").(*pop.Connection)
	s = c.Session()
	username := s.Get("username")
	if username != nil {
		a.Author = string(username.(string))
	} else {
		c.Set("message", "请先登录")
		return c.Render(422, r.HTML("message.html"))
	}
	uid := s.Get("uid")
	if uid != nil {
		a.Uid, _ = strconv.Atoi(string(s.Get("uid").(string)))
	} else {
		c.Set("message", "请先登录")
		return c.Render(422, r.HTML("message.html"))
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

	s := c.Session()
	c.Set("uid", s.Get("uid"))
	c.Set("username", s.Get("username"))
	c.Set("avatar", s.Get("avatar"))

	m := make([]map[string]string, 0)
	m = append(m, map[string]string{"gplus": "google"})
	c.Set("Providers", m)

	username := s.Get("username")
	if username == nil {
		c.Set("message", "请先登录")
		return c.Render(422, r.HTML("message.html"))
	}

	tid := c.Param("tid")
	if tid == "" {
		c.Set("message", "tid can not be empty")
		return c.Render(422, r.HTML("message.html"))
	}
	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)
	article := &models.Article{}
	// You can order your list here. Just change
	err := tx.Find(&article, tid)
	// to:
	// err := tx.Order("create_at desc").All(articles)
	if err != nil {
		return errors.WithStack(err)
	}
	// Make articles available inside the html template
	c.Set("article", article)

	return c.Render(200, r.HTML("article/edit.html"))
}
func ArticleSaveEdit(c buffalo.Context) error {

	s := c.Session()
	c.Set("uid", s.Get("uid"))
	c.Set("username", s.Get("username"))
	c.Set("avatar", s.Get("avatar"))

	m := make([]map[string]string, 0)
	m = append(m, map[string]string{"gplus": "google"})
	c.Set("Providers", m)

	username := s.Get("username")
	if username == nil {
		c.Set("message", "请先登录")
		return c.Render(422, r.HTML("message.html"))
	}

	tid := c.Param("tid")
	if tid == "" {
		c.Set("message", "tid can not be empty")
		return c.Render(422, r.HTML("message.html"))
	}
	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)
	article := &models.Article{}
	// You can order your list here. Just change
	err := tx.Find(&article, tid)
	// to:
	// err := tx.Order("create_at desc").All(articles)
	if err != nil {
		return errors.WithStack(err)
	}
	c.Request().ParseForm()
	article.Title = c.Request().Form.Get("title")
	article.Content = c.Request().Form.Get("content")

	// Validate the data from the html form
	verrs, err := tx.ValidateAndUpdate(article)
	if err != nil {
		return errors.WithStack(err)
	}
	if verrs.HasAny() {
		c.Set("errors", verrs)
		return c.Render(422, r.HTML("message.html"))
	}
	c.Set("message", "Update Success")

	return c.Render(200, r.JSON("message"))
}

// ArticleDelete default implementation.
func ArticleDelete(c buffalo.Context) error {

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
	err := tx.Order("created_at desc").All(articles)
	// to:
	// err := tx.Order("create_at desc").All(articles)
	if err != nil {
		return errors.WithStack(err)
	}
	// Make articles available inside the html template
	c.Set("articles", articles)

	return c.Render(200, r.HTML("article/delete.html"))
}
