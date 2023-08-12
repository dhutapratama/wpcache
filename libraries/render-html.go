package libraries

import (
	"fmt"
	"wpcache/models"

	"golang.org/x/net/html"
)

// Save content changes into new html version.
func RenderHtml(w models.Wordpress, n *html.Node) {
	fmt.Println("Loading: RenderHtml")
	fmt.Println()

	if w.RenderedIndex == nil {
		fmt.Println("Cant minify index")
		return
	}

	if err := html.Render(w.RenderedIndex, n); err != nil {
		fmt.Println(err)
		return
	}
}
