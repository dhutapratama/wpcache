package libraries

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"runtime"
	"strings"
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
	cache(w.Website, saveTo)
}

func cache(endPoint, cachePath string) (pathCache string) {
	var pathClean []string
	var path string
	pathDirty := strings.Split(cachePath, "/")

	// Build path
	lenPathDirty := len(pathDirty)
	lenPathDirectory := lenPathDirty - 1
	for i, v := range pathDirty {
		if v != "" && i > 0 {
			pathClean = append(pathClean, v)

			path = strings.Join(pathClean, "/")
			if runtime.GOOS == "linux" {
				path = fmt.Sprintf("/%s", path)
			} else if runtime.GOOS == "windows" {
				path = fmt.Sprintf("%s/%s", pathDirty[0], path)
			}

			if i < lenPathDirectory {
				if _, err := os.Stat(path); os.IsNotExist(err) {
					if err := os.Mkdir(path, 0755); err != nil {
						fmt.Println(err)
						return
					}
				}
			}
		}
	}

	if response, err := http.Get(endPoint); err != nil {
		fmt.Printf("%v\n", err)
		return
	} else {
		body, err := io.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("%v\n", err)
			return
		}

		fmt.Println("Cache Path: ", path)
		if err := os.WriteFile(path, body, 0644); err != nil {
			fmt.Println(err)
			return
		}
	}

	fmt.Println()
	return path
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

	parse_html(parsedIndexHtml, w.TempFolder, u, w.BundleCss)
}

func parse_html(n *html.Node, tempFolder string, u *url.URL, wCss io.Writer) {
	if n.Type == html.ElementNode {
		switch n.Data {
		case "link":
			parse_style(n, tempFolder, u, wCss)
		case "img":
		//	parse_img(n, tempFolder, u)
		case "script":
			//	parse_script(n, tempFolder, u)
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		parse_html(c, tempFolder, u, wCss)
	}
}
