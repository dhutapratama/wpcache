package main

import (
	"fmt"
	"wpcache/libraries"
)

func init() {
	libraries.LoadWordpressJson()
	libraries.VerifyPath()
	fmt.Println()
}

func main() {
	// using the function
	// mydir, err := os.Getwd()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(mydir)

	libraries.CachePage()
}