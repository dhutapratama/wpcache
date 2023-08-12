package libraries

import (
	"encoding/json"
	"fmt"
	"os"
	"wpcache/vars"
)

func LoadWordpressJson() {
	fmt.Println("Loading: WordpressJson")

	if configJson := os.Getenv("WPCACHE_JSON"); configJson == "" {
		fmt.Println("WPCACHE_JSON is not found in environment variables.")
		fmt.Println("Example Command 'export WPCACHE_JSON=/var/www/wordpress.json'")
	} else {
		vars.WordpressJson = configJson
	}

	byteData, err := os.ReadFile(vars.WordpressJson)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
		return
	}

	err = json.Unmarshal(byteData, &vars.Wordpress)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
		return
	}
}
