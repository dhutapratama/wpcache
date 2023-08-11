package libraries

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"wpcache/vars"
)

func minify_index(w io.Writer, fileIndex string) {
	var r io.Reader
	if f, err := os.OpenFile(fileIndex, os.O_RDONLY, 0644); err != nil {
		fmt.Println(err)
		return
	} else {
		r = bufio.NewReader(f)
		defer f.Close()
	}

	if err := vars.MinifierEngine.Minify("text/html", w, r); err != nil {
		fmt.Println(err)
		return
	}
}
