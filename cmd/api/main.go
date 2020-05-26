package main

import (
	"cybersport-comments-go/internal/config"
	"cybersport-comments-go/internal/server"
	"fmt"
	"log"
	"os"
)

func main() {
	config.Load()
	app := server.NewApp()
	if err := app.Run(config.Conf.Port); err != nil {
		log.Fatalf("%s", err.Error())
	}

	fmt.Printf("ENV PORT %s", os.Getenv("PORT"))
}
