package main

import (
	"flag"

	"github.com/maxon2034/trainee-go-tasks/loader/internal/app"
)

func main() {
	url := flag.String("url", "", "url")
	output := flag.String("output", "file", "downloaded file output")

	flag.Parse()

	app.Run(*url, *output)
}
