package models

import (
	"os"
)

type Wordpress struct {
	Name          string `json:"name"`
	Website       string `json:"website"`
	RootFolder    string `json:"root_folder"`
	TempFolder    string `json:"temp_folder"`
	SkipRenderJs  bool   `json:"skip_render_js"`
	SkipRenderCss bool   `json:"skip_render_css"`
	Verified      bool
	BundleCss     *os.File
	BundleJs      *os.File
	RenderedIndex *os.File
	MinifiedIndex *os.File
}
