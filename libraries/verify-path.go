package libraries

import (
	"fmt"
	"runtime"
	"strings"
	"wpcache/vars"
)

func VerifyPath() {
	fmt.Println("Loading: VerifyPath")

	for i, w := range vars.Wordpress {
		pathRoot := strings.Split(w.RootFolder, "/")
		pathTemp := strings.Split(w.TempFolder, "/")

		if len(pathRoot) < 2 || len(pathTemp) < 2 {
			continue
		}

		if runtime.GOOS == "windows" {
			driveRoot := strings.Split(pathRoot[0], ":")
			driveTemp := strings.Split(pathTemp[0], ":")

			if len(driveRoot) < 2 || len(driveTemp) < 2 {
				continue
			}
		} else if runtime.GOOS == "linux" {
			if pathRoot[0] != "" || pathTemp[0] != "" {
				continue
			}
		}

		vars.Wordpress[i].Verified = true
	}
}
