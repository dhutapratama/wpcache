package libraries

import (
	"fmt"
	"os"
	"path"
	"wpcache/helpers"
	"wpcache/vars"
)

func CreateBundleJs() {
	fmt.Println("Loading: Create Bundle Js")
	for i, w := range vars.Wordpress {

		bundleJs := path.Clean(fmt.Sprintf("%s/wp-cache/js/%s", w.RootFolder, "bundle.min.js"))
		if f, exist := helpers.CheckFile(bundleJs); exist {
			if f.IsDir() {
				fmt.Println("Cant render directory")
				return
			}
		}

		if err := os.WriteFile(bundleJs, []byte(""), 0644); err != nil {
			fmt.Println(err)
			return
		}

		if f, err := os.OpenFile(bundleJs, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644); err != nil {
			fmt.Println(err)
		} else {
			vars.Wordpress[i].BundleJs = f
		}
	}
}
