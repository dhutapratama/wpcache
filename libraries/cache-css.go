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

func parse_style(n *html.Node, w models.Wordpress, u *url.URL) {
	var rel, href string

	for _, element := range n.Attr {
		if element.Key == "rel" {
			rel = element.Val
		} else if element.Key == "href" {
			href = element.Val
		}
	}

	if rel == "stylesheet" {
		parse_css(href, w, u)
	}
}

func parse_css(href string, w models.Wordpress, u *url.URL) {
	uCss, err := url.Parse(href)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	if u.Host == uCss.Host {
		endPoint := fmt.Sprintf("%s://%s%s", uCss.Scheme, uCss.Host, uCss.EscapedPath())
		saveTo := fmt.Sprintf("%s/%s", w.TempFolder, uCss.EscapedPath())

		fmt.Printf("Caching Css: %s\n", href)
		fileCss := cache(endPoint, saveTo)

		minify_css(w, fileCss)
	}
}

func minify_css(w models.Wordpress, fileCss string) {
	var r io.Reader
	if f, err := os.OpenFile(fileCss, os.O_RDONLY, 0644); err != nil {
		fmt.Println(err)
		return
	} else {
		r = bufio.NewReader(f)
		defer f.Close()
	}

	if err := vars.MinifierEngine.Minify("text/css", w.BundleCss, r); err != nil {
		fmt.Println(err)
		return
	}
}
