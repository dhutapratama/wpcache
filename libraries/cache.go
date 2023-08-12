package libraries

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"wpcache/helpers"
	"wpcache/models"
)

// Download content and save to cache folder.
func cache(endPoint, escapedPath string, w models.Wordpress) (cache string) {

	cache = path.Clean(fmt.Sprintf("%s/%s", w.TempFolder, escapedPath))
	if IsDownloaded(endPoint) {
		return cache
	}

	// Check existing file
	if f, exist := helpers.CheckFile(cache); exist {
		if f.IsDir() {
			os.RemoveAll(cache)
		} else {
			fmt.Println("Content already downloaded")
			fmt.Println("Cache Path: ", cache)
			return
		}
	}

	// Download content and save to cachePath
	if response, err := http.Get(endPoint); err != nil {
		fmt.Printf("%v\n", err)
		return
	} else {
		body, err := io.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("%v\n", err)
			return
		}

		if err := os.WriteFile(cache, body, 0644); err != nil {
			fmt.Println(err)
			return
		}
	}

	fmt.Println()
	return cache
}

var DownloadedContent []string

func IsDownloaded(urlContent string) bool {
	for _, v := range DownloadedContent {
		if v == urlContent {
			return true
		}
	}

	DownloadedContent = append(DownloadedContent, urlContent)
	return false
}
