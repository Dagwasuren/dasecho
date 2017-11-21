package actions

import (
	"github.com/dasecho/dasecho/models"
	"github.com/gobuffalo/buffalo"
	"github.com/markbates/pop"
	"github.com/pkg/errors"
)

// TodaybestCreate default implementation.
func TodaybestCreate(c buffalo.Context) error {

	s := c.Session()
	c.Set("uid", s.Get("uid"))
	c.Set("username", s.Get("username"))
	c.Set("avatar", s.Get("avatar"))

	m := make([]map[string]string, 0)
	m = append(m, map[string]string{"gplus": "google"})
	c.Set("Providers", m)

	s = c.Session()
	username := s.Get("username")
	if username == nil {
		c.Set("message", "请先登录")
		return c.Render(422, r.HTML("message.html"))
	}
	return c.Render(200, r.HTML("todaybest/create.html"))
}

func TodaybestSaveCreate(c buffalo.Context) error {

	s := c.Session()
	c.Set("uid", s.Get("uid"))
	c.Set("username", s.Get("username"))
	c.Set("avatar", s.Get("avatar"))

	m := make([]map[string]string, 0)
	m = append(m, map[string]string{"gplus": "google"})
	c.Set("Providers", m)

	tx := c.Value("tx").(*pop.Connection)

	a := &models.Todaybest{}
	c.Request().ParseForm()
	a.Content = c.Request().Form.Get("content")
	// Get the DB connection from the context
	s = c.Session()
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

// TodaybestEdit default implementation.
func TodaybestEdit(c buffalo.Context) error {

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
	todaybest := &models.Todaybest{}
	// You can order your list here. Just change
	err := tx.Where("id = ?", tid).First(todaybest)
	// to:
	// err := tx.Order("create_at desc").All(Todaybests)
	if err != nil {
		return errors.WithStack(err)
	}
	c.Set("todaybest", todaybest)

	return c.Render(200, r.HTML("todaybest/edit.html"))
}

func TodaybestList(c buffalo.Context) error {

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

	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)
	todaybests := &models.Todaybests{}

	err := tx.Order("created_at desc").All(todaybests)
	// to:
	// err := tx.Order("create_at desc").All(articles)
	if err != nil {
		return errors.WithStack(err)
	}
	// Make articles available inside the html template
	c.Set("todaybests", todaybests)

	return c.Render(200, r.HTML("todaybest/list.html"))
}
func TodaybestSaveEdit(c buffalo.Context) error {
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

	c.Request().ParseForm()
	tid := c.Request().Form.Get("tid")
	if tid == "" {
		c.Set("message", "tid can not be empty")
		return c.Render(422, r.HTML("message.html"))
	}
	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)
	Todaybest := &models.Todaybest{}
	// You can order your list here. Just change

	err := tx.Where("id = ?", tid).First(Todaybest)
	// to:
	// err := tx.Order("create_at desc").All(Todaybests)
	if err != nil {
		return errors.WithStack(err)
	}
	Todaybest.Content = c.Request().Form.Get("content")

	// Validate the data from the html form
	verrs, err := tx.ValidateAndUpdate(Todaybest)
	if err != nil {
		return errors.WithStack(err)
	}
	if verrs.HasAny() {
		c.Set("errors", verrs)
		return c.Render(422, r.HTML("message.html"))
	}
	c.Set("message", "Update Success")

	return c.Render(200, r.HTML("message.html"))
}

// TodaybestDelete default implementation.
func TodaybestDelete(c buffalo.Context) error {

	s := c.Session()
	c.Set("uid", s.Get("uid"))
	c.Set("username", s.Get("username"))
	c.Set("avatar", s.Get("avatar"))

	m := make([]map[string]string, 0)
	m = append(m, map[string]string{"gplus": "google"})
	c.Set("Providers", m)

	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)
	Todaybests := &models.Todaybests{}
	// You can order your list here. Just change
	err := tx.Order("created_at desc").All(Todaybests)
	// to:
	// err := tx.Order("create_at desc").All(Todaybests)
	if err != nil {
		return errors.WithStack(err)
	}
	// Make Todaybests available inside the html template
	c.Set("Todaybests", Todaybests)

	return c.Render(200, r.HTML("todaybest/delete.html"))
}
