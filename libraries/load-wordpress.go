package libraries

import (
	"encoding/json"
	"fmt"
	"os"
	"wpcache/vars"
)

func LoadWordpressJson() {
	fmt.Println("Loading: WordpressJson")
	byteData, err := os.ReadFile(vars.WordpressJson)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = json.Unmarshal(byteData, &vars.Wordpress)
	if err != nil {
		fmt.Println(err)
		return
	}
}
