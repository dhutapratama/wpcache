package libraries

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/nickalie/go-webpbin"
)

func ProcessWebp(imgPath string) {
	// file, err := os.Open(imgPath)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// img, err := jpeg.Decode(file)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// Slice path
	filePath := strings.Split(imgPath, "/")
	imgFile := filePath[len(filePath)-1]
	imgName := imgFile[:len(imgFile)-len(filepath.Ext(imgFile))]
	imgWebp := fmt.Sprintf("%s.webp", imgName)
	filePath[len(filePath)-1] = imgWebp

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

	err := webpbin.NewCWebP().
		Quality(80).
		InputFile(imgPath).
		OutputFile(strings.Join(filePath, "/")).
		Run()
	if err != nil {
		fmt.Println(err)
		return
	}
}
