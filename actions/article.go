package actions

import "github.com/gobuffalo/buffalo"

// ArticleCreate default implementation.
func ArticleCreate(c buffalo.Context) error {
	return c.Render(200, r.HTML("article/create.html"))
}

func ArticleSaveCreate(c buffalo.Context) error {
	return c.Render(200, r.JSON("{\"hello\":\"message\"}"))
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
