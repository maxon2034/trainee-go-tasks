package main

import (
	"context"
	"log"
	"log/slog"
	"os"

	"github.com/maxon2034/trainee-go-tasks/compositehandler/internal/app"
)

func main() {
	ctx := context.Background()

	file, err := os.Create("log.json")
	if err != nil {
		log.Fatal("Error in creating file: ", err)
	}

	JSONHandler := slog.NewJSONHandler(file, nil)
	TextHandler := slog.NewTextHandler(os.Stdout, nil)

	app.Run(ctx, JSONHandler, TextHandler)
}
