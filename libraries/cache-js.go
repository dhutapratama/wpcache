package libraries

import (
	"bufio"
	"fmt"
	"io"
	"net/url"
	"os"
	"wpcache/models"
	"wpcache/vars"

	"golang.org/x/net/html"
)

func parse_script(n *html.Node, w models.Wordpress, u *url.URL) {
	for _, element := range n.Attr {
		if element.Key == "src" {

			uJs, err := url.Parse(element.Val)
			if err != nil {
				fmt.Printf("%s\n", err)
				return
			}

			if u.Host == uJs.Host {
				endPoint := fmt.Sprintf("%s://%s%s", uJs.Scheme, uJs.Host, uJs.EscapedPath())
				saveTo := fmt.Sprintf("%s/%s", w.TempFolder, uJs.EscapedPath())

				fmt.Printf("Caching Script: %s\n", element.Val)
				fileJs := cache(endPoint, saveTo)

				minify_js(w, fileJs)
			}
		}
	}
}

func minify_js(w models.Wordpress, fileJs string) {
	var r io.Reader
	if f, err := os.OpenFile(fileJs, os.O_RDONLY, 0644); err != nil {
		fmt.Println(err)
		return
	} else {
		r = bufio.NewReader(f)
		defer f.Close()
	}

	if err := vars.MinifierEngine.Minify("application/javascript", w.BundleJs, r); err != nil {
		fmt.Println(err)
		return
	}
}
