package main

import (
	"github.com/JaiiR320/carlistingsaver/api"
)

func main() {
	app := api.NewServer(":3000")
	app.Run()
}
