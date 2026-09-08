package app

import (
	"context"
	"fmt"
	"time"

	"github.com/maxon2034/trainee-go-tasks/bakery/internal/bakery"
)

func Run() {
	ctx := context.Background()
	start := time.Now()
	defer func() {
		finish := time.Since(start)
		fmt.Println("finished in ", finish.Seconds(), " seconds")
	}()
	bakery.Start(ctx, 1000, 10, 20, 1, 2)
}
