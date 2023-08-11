package libraries

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"wpcache/vars"
)

func CreateMinifiedIndex() {
	fmt.Println("Loading: Create Minified Index")
	for i, w := range vars.Wordpress {
		bundleFile := "index.html"
		saveDirDirty := fmt.Sprintf("%s/wp-cache/", w.RootFolder)

		pathDirty := strings.Split(saveDirDirty, "/")
		var pathClean []string
		var path string

		// Build: Directory Path
		for i, v := range pathDirty {
			if v != "" && i > 0 {
				pathClean = append(pathClean, v)

				path = strings.Join(pathClean, "/")

				if runtime.GOOS == "linux" {
					path = fmt.Sprintf("/%s", path)
				} else if runtime.GOOS == "windows" {
					path = fmt.Sprintf("%s/%s", pathDirty[0], path)
				}

				if _, err := os.Stat(path); os.IsNotExist(err) {
					if err := os.Mkdir(path, 0755); err != nil {
						fmt.Println(err)
						return
					}
				}
			}
		}

		// Build: File Path
		minifiedIndex := fmt.Sprintf("%s/%s", path, bundleFile)
		fmt.Println(minifiedIndex)

		if err := os.WriteFile(minifiedIndex, []byte(""), 0644); err != nil {
			fmt.Println(err)
			return
		}

		if f, err := os.OpenFile(minifiedIndex, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644); err != nil {
			fmt.Println(err)
		} else {
			vars.Wordpress[i].MinifiedIndex = f
		}
	}
}
