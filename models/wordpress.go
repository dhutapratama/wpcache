package models

import (
	"os"
)

type Wordpress struct {
	Name          string `json:"name"`
	Website       string `json:"website"`
	RootFolder    string `json:"root_folder"`
	TempFolder    string `json:"temp_folder"`
	Verified      bool
	BundleCss     *os.File
	BundleJs      *os.File
	MinifiedIndex *os.File
}
