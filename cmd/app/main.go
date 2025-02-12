package main

import (
	"context"
	"github.com/Rasikrr/learning_platform/internal/app"
	"log"
)

const (
	appName = "learning_platform"
)

func main() {
	ctx := context.Background()
	a := app.InitApp(ctx, appName)

	if err := a.Start(ctx); err != nil {
		log.Fatal(err)
	}
}
