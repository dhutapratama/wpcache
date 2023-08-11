package libraries

import (
	"fmt"
	"net/url"

	"golang.org/x/net/html"
)

func parse_script(n *html.Node, tempFolder string, u *url.URL) {
	for _, element := range n.Attr {
		if element.Key == "src" {

			uJs, err := url.Parse(element.Val)
			if err != nil {
				fmt.Printf("%s\n", err)
				return
			}

			if u.Host == uJs.Host {
				endPoint := fmt.Sprintf("%s://%s%s", uJs.Scheme, uJs.Host, uJs.EscapedPath())
				saveTo := fmt.Sprintf("%s/%s", tempFolder, uJs.EscapedPath())

				fmt.Printf("Caching Script: %s\n", element.Val)
				cache(endPoint, saveTo)
			}
		}
	}
}
