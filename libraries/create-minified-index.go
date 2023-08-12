package libraries

import (
	"fmt"
	"os"
	"path"
	"wpcache/helpers"
	"wpcache/vars"
)

func CreateMinifiedIndex() {
	fmt.Println("Loading: Create Minified Index")
	for i, w := range vars.Wordpress {

		indexHtml := path.Clean(fmt.Sprintf("%s/%s", w.RootFolder, "index.html"))
		if f, exist := helpers.CheckFile(indexHtml); exist {
			if f.IsDir() {
				fmt.Println("Cant render directory")
				return
			}
		}

		// Build: File Path
		if err := os.WriteFile(indexHtml, []byte(""), 0644); err != nil {
			fmt.Println(err)
			return
		}

		if f, err := os.OpenFile(indexHtml, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644); err != nil {
			fmt.Println(err)
		} else {
			vars.Wordpress[i].MinifiedIndex = f
		}
	}
}
