package libraries

import (
	"fmt"
	"net/url"

	"golang.org/x/net/html"
)

func parse_img(n *html.Node, tempFolder string, u *url.URL) {
	for _, element := range n.Attr {
		if element.Key == "src" {
			uImg, err := url.Parse(element.Val)
			if err != nil {
				fmt.Printf("%s\n", err)
				return
			}

			if u.Host == uImg.Host {
				endPoint := fmt.Sprintf("%s://%s%s", uImg.Scheme, uImg.Host, uImg.EscapedPath())
				saveTo := fmt.Sprintf("%s/%s", tempFolder, uImg.EscapedPath())

				fmt.Printf("Caching Img: %s\n", element.Val)
				imgPath := cache(endPoint, saveTo)
				ProcessWebp(imgPath)
			}
		}
	}
}
