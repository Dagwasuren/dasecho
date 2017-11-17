package actions

import (
	"fmt"
	"os"
	"log"
	"github.com/gobuffalo/buffalo"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
  "github.com/markbates/goth/providers/gplus"

)

func init() {
	gothic.Store = App().SessionStore
	log.Println("The app Host: ", app.Host)
	goth.UseProviders(
		gplus.New(os.Getenv("GPLUS_KEY"), os.Getenv("GPLUS_SECRET"), fmt.Sprintf("%s%s", app.Host, "/auth/gplus/callback")),
	)
}


func AuthCallback(c buffalo.Context) error {
	user, err := gothic.CompleteUserAuth(c.Response(), c.Request())
	if err != nil {
		return c.Error(401, err)
	}
	// Do something with the user, maybe register them/sign them in
	return c.Render(200, r.JSON(user))
}
