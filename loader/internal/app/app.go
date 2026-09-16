package app

import (
	"log"

	"github.com/maxon2034/trainee-go-tasks/loader/internal/loader"
)

func Run(fileURL, fileName string) {
	err := loader.DownloadFile(fileURL, fileName)
	if err != nil {
		log.Fatal(err)
	}
}
