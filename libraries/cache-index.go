package libraries

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path"
	"wpcache/models"
	"wpcache/vars"
)

func RenderMinifiedHtml(w models.Wordpress) {
	fileIndex := path.Clean(fmt.Sprintf("%s/%s", w.TempFolder, "index_rendered.html"))

	var r io.Reader
	if f, err := os.OpenFile(fileIndex, os.O_RDONLY, 0644); err != nil {
		fmt.Println(err)
		return
	} else {
		r = bufio.NewReader(f)
		defer f.Close()
	}

	if err := vars.MinifierEngine.Minify("text/html", w.MinifiedIndex, r); err != nil {
		fmt.Println(err)
		return
	}
}
