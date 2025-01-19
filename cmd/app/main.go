package main

import (
	"context"
	"github.com/Rasikrr/learning_platform/internal/app"
	"log"
)

const (
	appName = "assigment_2"
	c       = "0123456789"
)

func main() {
	ctx := context.Background()
	a := app.InitApp(ctx, appName)
	if err := a.Start(ctx); err != nil {
		log.Fatal(err)
	}
}
