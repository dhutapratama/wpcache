package libraries

import (
	"fmt"
	"os"
	"path"
	"wpcache/helpers"
	"wpcache/vars"
)

func CreateBundleCss() {
	fmt.Println("Loading: Create Bundle CSS")
	for i, w := range vars.Wordpress {
		bundleCss := path.Clean(fmt.Sprintf("%s/wp-cache/css/%s", w.RootFolder, "bundle.min.css"))
		helpers.CheckFile(bundleCss)

		// Build: File Path
		if err := os.WriteFile(bundleCss, []byte(""), 0644); err != nil {
			fmt.Println(err)
			return
		}

		if f, err := os.OpenFile(bundleCss, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644); err != nil {
			fmt.Println(err)
		} else {
			vars.Wordpress[i].BundleCss = f
		}
	}
}
