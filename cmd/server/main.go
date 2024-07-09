package main

import (
	"github.com/luizgr/go-expert-api/configs"
)

func main() {
	config, _ := configs.LoadConfig(".")
	println(config.DBHost)
}
