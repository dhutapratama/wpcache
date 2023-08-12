package helpers

import (
	"fmt"
	"io/fs"
	"os"
	"path"
	"strings"
)

var Exist []string

// Check existing directory
func CheckDir(dirPath string) (fileinfo fs.FileInfo) {
	dirPath = path.Clean(dirPath)

	if f, err := os.Stat(dirPath); !os.IsNotExist(err) {
		if f.IsDir() {
			return f
		}
	}

	dirPathArr := strings.Split(dirPath, "/")
	for i := range dirPathArr {
		if i > 0 {
			dirCurrent := path.Join(dirPathArr[:i+1]...)
			// fmt.Println("Current ", dirCurrent)
			// if runtime.GOOS == "linux" {
			// 	pathOutput = fmt.Sprintf("/%s", pathOutput)
			// } else if runtime.GOOS == "windows" {
			// 	pathOutput = fmt.Sprintf("%s/%s", pathDirty[0], pathOutput)
			// }

			if _, err := os.Stat(dirCurrent); os.IsNotExist(err) {
				if err := os.Mkdir(dirCurrent, 0755); err != nil {
					fmt.Println(err)
					return
				}
			}
		}
	}
	return CheckDir(dirPath)
}

// Check existing file
func CheckFile(filePath string) (fileInfo fs.FileInfo, isExist bool) {
	filePath = path.Clean(filePath)
	if f, err := os.Stat(filePath); os.IsNotExist(err) {
		// Now we make sure the folder is exist
		dirPath := path.Dir(filePath)

		CheckDir(dirPath)

		return nil, false
	} else {
		return f, true
	}
}
