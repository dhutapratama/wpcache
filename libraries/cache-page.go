package libraries

import (
	"bytes"
	"fmt"
	"net/url"
	"os"
	"path"
	"time"
	"wpcache/helpers"
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
		RenderHtml(w, GetAssets(w))
		RenderMinifiedHtml(w)
	}
}

func GetPage(w models.Wordpress) {
	fmt.Println("Fetching HTML: ", w.Website)

	u, err := url.Parse(w.Website)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	indexHtml := path.Clean(fmt.Sprintf("%s/%s", w.TempFolder, "index.html"))

	// Check existing file
	if f, exist := helpers.CheckFile(indexHtml); exist {
		if f.IsDir() {
			os.RemoveAll(indexHtml)
		} else {
			os.Remove(indexHtml)
		}
	}

	cache(fmt.Sprintf("%s://%s%s?cache=false", u.Scheme, u.Host, u.EscapedPath()), "index.html", w)
}

func GetAssets(w models.Wordpress) *html.Node {
	fmt.Println("Fetching Assets: ", w.Website)
	fmt.Println()

	readTo := fmt.Sprintf("%s/%s", w.TempFolder, "index.html")
	indexHtml, err := os.ReadFile(readTo)
	if err != nil {
		fmt.Printf("%s\n", err)
		return nil
	}

	parsedIndexHtml, err := html.Parse(bytes.NewReader(indexHtml))
	if err != nil {
		fmt.Printf("%v\n", err)
		return nil
	}

	u, err := url.Parse(w.Website)
	if err != nil {
		fmt.Printf("%v\n", err)
		return parsedIndexHtml
	}

	parse_html(parsedIndexHtml, w, u)
	append_bundle(parsedIndexHtml, u)

	return parsedIndexHtml
}

func parse_html(n *html.Node, w models.Wordpress, u *url.URL) (remove *html.Node) {
	var removechild []*html.Node

	if n.Type == html.ElementNode {
		switch n.Data {
		case "link":
			if parse_style(n, w, u) {
				remove = n
			}
		case "img":
			parse_img(n, w, u)
		case "script":
			if parse_script(n, w, u) {
				remove = n
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if toRemove := parse_html(c, w, u); toRemove != nil {
			removechild = append(removechild, toRemove)
		}
	}

	for _, child := range removechild {
		n.RemoveChild(child)
	}

	return
}

func append_bundle(n *html.Node, u *url.URL) {
	hash := time.Now().Format("20060102150405")

	if n.Type == html.ElementNode {
		switch n.Data {
		// Append Bundle Css
		case "head":
			n.AppendChild(&html.Node{
				Type: html.ElementNode,
				Data: "link",
				Attr: []html.Attribute{
					{
						Key: "rel",
						Val: "stylesheet",
					},
					{
						Key: "id",
						Val: "cache-bundle-css",
					},
					{
						Key: "href",
						Val: fmt.Sprintf("%s://%s%s?v=%s", u.Scheme, u.Host, "/wp-cache/css/bundle.min.css", hash),
					},
					{
						Key: "media",
						Val: "all",
					},
				},
			})
		// Append Bundle Js
		case "body":
			n.AppendChild(&html.Node{
				Type: html.ElementNode,
				Data: "script",
				Attr: []html.Attribute{
					{
						Key: "id",
						Val: "cache-bundle-js",
					},
					{
						Key: "src",
						Val: fmt.Sprintf("%s://%s%s?v=%s", u.Scheme, u.Host, "/wp-cache/js/bundle.min.js", hash),
					},
				},
			})
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		append_bundle(c, u)
	}
}
