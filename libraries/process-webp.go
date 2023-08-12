package libraries

import (
	"fmt"
	"os"
	"path"
	"strings"
	"wpcache/helpers"
	"wpcache/models"

	"github.com/nickalie/go-webpbin"
)

func ProcessWebp(imgPath string, escapedPath string, w models.Wordpress) (cache string, newEscPath string) {

	// saveDirDirty := fmt.Sprintf("%s/wp-cache/%s", w.RootFolder, orig)
	// pathDirty := strings.Split(saveDirDirty, "/")
	// var pathClean []string
	// var path string

	// Slice path
	// filePath := strings.Split(escapedPath, "/")
	// fileExt :=
	newEscPath = fmt.Sprintf("/wp-cache%s.webp", strings.TrimSuffix(escapedPath, path.Ext(escapedPath)))

	// imgFile := filePath[len(filePath)-1]
	// imgName := imgFile[:len(imgFile)-len(filepath.Ext(imgFile))]
	// imgWebp := fmt.Sprintf("%s.webp", imgName)
	// filePath[len(filePath)-1] = imgWebp
	// filePath := strings.Split(cache, "/")

	// Check existing file
	cache = path.Clean(fmt.Sprintf("%s%s", w.RootFolder, newEscPath))
	if f, exist := helpers.CheckFile(cache); exist {
		if f.IsDir() {
			os.RemoveAll(cache)
		} else {
			fmt.Println("Content already webp processed")
			fmt.Println("Webp Path: ", cache)
			fmt.Println()
			return
		}
	}

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

	// // Slice path
	// filePath := strings.Split(orig, "/")
	// imgFile := filePath[len(filePath)-1]
	// imgName := imgFile[:len(imgFile)-len(filepath.Ext(imgFile))]
	// imgWebp := fmt.Sprintf("%s.webp", imgName)
	// // filePath[len(filePath)-1] = imgWebp

	// output := fmt.Sprintf("%s/%s", path, imgWebp)

	err := webpbin.NewCWebP().
		Quality(80).
		InputFile(imgPath).
		OutputFile(cache).
		Run()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Webp Path: ", cache)
	fmt.Println()

	return
}
