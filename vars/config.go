package vars

import (
	"wpcache/models"

	"github.com/tdewolff/minify/v2"
)

var (
	WordpressJson  string = "wordpress.json"
	Wordpress      []models.Wordpress
	MinifierEngine *minify.M
)
