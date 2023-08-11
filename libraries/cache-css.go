package libraries

import (
	"bufio"
	"fmt"
	"io"
	"net/url"
	"os"
	"wpcache/vars"

	"golang.org/x/net/html"
)

func parse_style(n *html.Node, tempFolder string, u *url.URL, w io.Writer) {
	var rel, href string

	for _, element := range n.Attr {
		if element.Key == "rel" {
			rel = element.Val
		} else if element.Key == "href" {
			href = element.Val
		}
	}

	if rel == "stylesheet" {
		parse_css(href, tempFolder, u, w)
	}
}

func parse_css(href, tempFolder string, u *url.URL, w io.Writer) {
	uCss, err := url.Parse(href)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	if u.Host == uCss.Host {
		endPoint := fmt.Sprintf("%s://%s%s", uCss.Scheme, uCss.Host, uCss.EscapedPath())
		saveTo := fmt.Sprintf("%s/%s", tempFolder, uCss.EscapedPath())

		fmt.Printf("Caching Css: %s\n", href)
		fileCss := cache(endPoint, saveTo)

		minify_css(w, fileCss)
	}
}

func minify_css(w io.Writer, fileCss string) {
	var r io.Reader
	if f, err := os.OpenFile(fileCss, os.O_RDONLY, 0644); err != nil {
		fmt.Println(err)
		return
	} else {
		r = bufio.NewReader(f)
		defer f.Close()
	}

	if err := vars.MinifierEngine.Minify("text/css", w, r); err != nil {
		fmt.Println(err)
		return
	}
}
