package helpers

// import (
// 	"fmt"
// 	"os"
// 	"runtime"
// 	"strings"
// )

// // Cleaning Path from invalid directory
// // and create directory if not exist
// func CleanPath(path string) string {
// 	var pathClean []string
// 	var pathOutput string

// 	pathDirty := strings.Split(path, "/")
// 	lenPathDirty := len(pathDirty)
// 	lenPathDirtyDir := lenPathDirty - 1

// 	// Rebuild clean path
// 	for i, v := range pathDirty {
// 		if v != "" && i > 0 {
// 			pathClean = append(pathClean, v)

// 			// Rejoin path
// 			pathOutput = strings.Join(pathClean, "/")
// 			if runtime.GOOS == "linux" {
// 				pathOutput = fmt.Sprintf("/%s", path)
// 			} else if runtime.GOOS == "windows" {
// 				pathOutput = fmt.Sprintf("%s/%s", pathDirty[0], path)
// 			}

// 			// Check existing directory
// 			if i < lenPathDirtyDir {
// 				if _, err := os.Stat(path); os.IsNotExist(err) {
// 					if err := os.Mkdir(path, 0755); err != nil {
// 						fmt.Println(err)
// 						return path
// 					}
// 				}
// 			}
// 		}
// 	}

// 	return pathOutput
// }
