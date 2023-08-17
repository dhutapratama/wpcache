package libraries

import (
	"crypto/rand"
	"fmt"
	"os"
	"path"
	"wpcache/helpers"
	"wpcache/models"
)

func cache_string(content []byte, w models.Wordpress, extension string) (cache string) {
	for {
		b := make([]byte, 8)
		rand.Read(b)

		cache = path.Clean(fmt.Sprintf("%s/%x%s", w.TempFolder, b, extension))
		// Check existing file
		if f, exist := helpers.CheckFile(cache); exist {
			if f.IsDir() {
				os.RemoveAll(cache)
			}
		} else {
			break
		}
	}

	if err := os.WriteFile(cache, content, 0644); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println()
	return cache
}
