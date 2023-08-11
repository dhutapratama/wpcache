package libraries

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"
	"strings"
)

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
