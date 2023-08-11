package libraries

import (
	"bytes"
	"fmt"
	"io"
	"net/url"
	"os"
	"wpcache/models"
	"wpcache/vars"

	"golang.org/x/net/html"
)

func CachePage() {
	fmt.Println("Loading: CachePage")
	fmt.Println("")

	for _, w := range vars.Wordpress {
		if !w.Verified {
			fmt.Println("Invalid path format: ", w.Name)
		}

		fmt.Println("Processing: ", w.Name)

		GetPage(w)
		GetAssets(w)
	}
}

func GetPage(w models.Wordpress) {
	fmt.Println("Fetching HTML: ", w.Website)

	saveTo := fmt.Sprintf("%s/%s", w.TempFolder, "index.html")
	fileIndex := cache(w.Website, saveTo)
	minify_index(w.MinifiedIndex, fileIndex)
}

func GetAssets(w models.Wordpress) {
	fmt.Println("Fetching Assets: ", w.Website)
	fmt.Println()

	readTo := fmt.Sprintf("%s/%s", w.TempFolder, "index.html")
	indexHtml, err := os.ReadFile(readTo)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	parsedIndexHtml, err := html.Parse(bytes.NewReader(indexHtml))
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	u, err := url.Parse(w.Website)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	parse_html(parsedIndexHtml, w.TempFolder, u, w.BundleCss, w.BundleJs)
}

func parse_html(n *html.Node, tempFolder string, u *url.URL, wCss io.Writer, wJs io.Writer) {
	if n.Type == html.ElementNode {
		switch n.Data {
		case "link":
			parse_style(n, tempFolder, u, wCss)
		case "img":
			parse_img(n, tempFolder, u)
		case "script":
			parse_script(n, tempFolder, u, wJs)
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		parse_html(c, tempFolder, u, wCss, wJs)
	}
}
