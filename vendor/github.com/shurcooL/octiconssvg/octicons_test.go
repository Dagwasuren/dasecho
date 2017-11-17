package octiconssvg_test

import (
	"io"
	"log"
	"os"

	"github.com/shurcooL/octiconssvg"
	"golang.org/x/net/html"
)

func Example() {
	var w io.Writer = os.Stdout // Or, e.g., http.ResponseWriter in your handler, etc.

	err := html.Render(w, octiconssvg.Alert())
	if err != nil {
		log.Fatalln(err)
	}

	// Output:
	// <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 16 16" style="fill: currentColor; vertical-align: top;"><path d="M8.865 1.52c-.18-.31-.51-.5-.87-.5s-.69.19-.87.5L.275 13.5c-.18.31-.18.69 0 1 .19.31.52.5.87.5h13.7c.36 0 .69-.19.86-.5.17-.31.18-.69.01-1L8.865 1.52zM8.995 13h-2v-2h2v2zm0-3h-2V6h2v4z"></path></svg>
}
