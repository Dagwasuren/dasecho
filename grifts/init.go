package grifts

import (
	"github.com/dasecho/dasecho/actions"
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
