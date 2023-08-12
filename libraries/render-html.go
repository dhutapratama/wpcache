package libraries

import (
	"fmt"
	"os"
	"path"
	"wpcache/models"

	"golang.org/x/net/html"
)

// Save content changes into new html version.
func RenderHtml(w models.Wordpress, n *html.Node) {
	fmt.Println("Loading: RenderHtml")
	fmt.Println()

	saveTo := fmt.Sprintf("%s/%s", w.TempFolder, "index_rendered.html")
	saveTo = path.Clean(saveTo)

	if file, err := os.OpenFile(saveTo, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0644); err != nil {
		fmt.Println(err)
		return
	} else {
		html.Render(file, n)
	}
}
