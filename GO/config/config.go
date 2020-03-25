package config

import (
	"os"
)


var SECRETKEY []byte


func Load() {
	SECRETKEY = []byte(os.Getenv("API_SECRET"))
}

