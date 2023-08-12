package libraries

import (
	"fmt"
	"net/url"
	"path"
	"strings"
	"wpcache/models"

	"golang.org/x/net/html"
)

// Parse the image value
func parse_img(n *html.Node, w models.Wordpress, u *url.URL) {
	for i, element := range n.Attr {
		if element.Key == "src" {
			fmt.Println("Caching Img: ", element.Val)

			uImg, err := url.Parse(element.Val)
			if err != nil {
				fmt.Printf("%s\n", err)
				return
			}

			// Check if already processed
			if IsWebp(uImg.EscapedPath()) {
				fmt.Println("Already webp, ignoring")
				fmt.Println()
				continue
			}

			// Check origin
			if u.Host == uImg.Host {
				imgPath := cache(element.Val, uImg.EscapedPath(), w)
				_, escPath := ProcessWebp(imgPath, uImg.EscapedPath(), w)

				n.Attr[i].Val = fmt.Sprintf("%s://%s%s", uImg.Scheme, uImg.Host, escPath)
			} else {
				fmt.Println("Outside origin ignoring")
				fmt.Println()
			}
		} else if element.Key == "srcset" {
			var srcsetUrl, srcsetOpt, srcset []string

			valAntiComma := strings.Split(element.Val, ",")
			for _, v := range valAntiComma {
				valAntiSpace := strings.Split(v, " ")
				for _, v2 := range valAntiSpace {

					// Skip undefined
					if v2 == "" {
						continue
					}

					uImg, err := url.ParseRequestURI(v2)
					if err != nil {
						srcsetOpt = append(srcsetOpt, v2)

						// fmt.Printf("%s\n", err)
						fmt.Println()
						continue
					}

					fmt.Println("Caching Img: ", v2)

					// Check if already processed
					if IsWebp(uImg.EscapedPath()) {
						fmt.Println("Already webp, ignoring")
						fmt.Println()
						continue
					}

					if u.Host == uImg.Host {
						imgPath := cache(v2, uImg.EscapedPath(), w)

						_, urlWebp := ProcessWebp(imgPath, uImg.EscapedPath(), w)
						srcsetUrl = append(srcsetUrl, urlWebp)
					} else {
						fmt.Println("Outside origin ignoring")
						fmt.Println()
					}
				}
			}

			// Rebuild Srcset
			for i, v := range srcsetUrl {
				srcset = append(srcset, fmt.Sprintf("%s %s", v, srcsetOpt[i]))
			}
			n.Attr[i].Val = strings.Join(srcset, ", ")
		}
	}
}

func IsWebp(escapedPath string) bool {
	return path.Ext(escapedPath) == ".webp"
}
