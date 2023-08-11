package libraries

import (
	"fmt"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func parse_img(n *html.Node, tempFolder string, u *url.URL) {
	var urls []string

	for _, element := range n.Attr {
		if element.Key == "src" {
			uImg, err := url.Parse(element.Val)
			if err != nil {
				fmt.Printf("%s\n", err)
				return
			}

			// Check if already processed
			ok := true
			for _, v := range urls {
				if v == element.Val {
					ok = false
					break
				}
			}
			if !ok {
				fmt.Println("Already processed: ", element.Val)
				fmt.Println()
				continue
			}
			urls = append(urls, element.Val)

			if u.Host == uImg.Host {
				endPoint := fmt.Sprintf("%s://%s%s", uImg.Scheme, uImg.Host, uImg.EscapedPath())
				saveTo := fmt.Sprintf("%s/%s", tempFolder, uImg.EscapedPath())

				fmt.Printf("Caching Img: %s\n", element.Val)
				imgPath := cache(endPoint, saveTo)
				ProcessWebp(imgPath)
			}
		} else if element.Key == "srcset" {
			valAntiComma := strings.Split(element.Val, ",")
			for _, v := range valAntiComma {
				valAntiSpace := strings.Split(v, " ")
				for _, v2 := range valAntiSpace {

					// Skip undefined
					if v2 == "" {
						continue
					}

					uImg, err := url.Parse(v2)
					if err != nil {
						fmt.Printf("%s\n", err)
						continue
					}

					// Check if already processed
					ok := true
					for _, v3 := range urls {
						if v3 == v2 {
							ok = false
							break
						}
					}
					if !ok {
						fmt.Println("Already processed: ", v2)
						fmt.Println()
						continue
					}
					urls = append(urls, v2)

					if u.Host == uImg.Host {
						endPoint := fmt.Sprintf("%s://%s%s", uImg.Scheme, uImg.Host, uImg.EscapedPath())
						saveTo := fmt.Sprintf("%s/%s", tempFolder, uImg.EscapedPath())

						fmt.Printf("Caching Img: %s\n", v2)
						imgPath := cache(endPoint, saveTo)
						ProcessWebp(imgPath)
					}
				}
			}
		}
	}
}
