package libraries

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"wpcache/models"

	"github.com/nickalie/go-webpbin"
)

func ProcessWebp(imgPath string, orig string, w models.Wordpress) {
	// file, err := os.Open(imgPath)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// img, err := jpeg.Decode(file)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// output, err := os.Create(filepath.Join("/"))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer output.Close()

	// options, err := encoder.NewLossyEncoderOptions(encoder.PresetDefault, 75)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// if err := webp.Encode(output, img, options); err != nil {
	// 	log.Fatalln(err)
	// }

	saveDirDirty := fmt.Sprintf("%s/wp-cache/%s", w.RootFolder, orig)
	pathDirty := strings.Split(saveDirDirty, "/")
	var pathClean []string
	var path string

	// Build: Directory Path
	for i, v := range pathDirty[:len(pathDirty)-1] {
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

	// Slice path
	filePath := strings.Split(orig, "/")
	imgFile := filePath[len(filePath)-1]
	imgName := imgFile[:len(imgFile)-len(filepath.Ext(imgFile))]
	imgWebp := fmt.Sprintf("%s.webp", imgName)
	// filePath[len(filePath)-1] = imgWebp

	output := fmt.Sprintf("%s/%s", path, imgWebp)
	fmt.Println("Processed Path: ", output)

	err := webpbin.NewCWebP().
		Quality(80).
		InputFile(imgPath).
		OutputFile(output).
		Run()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println()
}
