package main

import (
	"context"
	"log"

	"github.com/go/mini_market/src/config"
	"github.com/go/mini_market/src/routers"
	"github.com/go/mini_market/src/server"
)

func main() {
	config := config.GetVariablesEnv()

	s, err := server.NewServer(context.Background(), &server.Config{
		Port:        config.Port,
		DatabaseUrl: config.DatabaseUrl,
		JWTSecret:   config.JWTSecret,
	})
	if err != nil {
		log.Fatal(err)
	}

	s.Start(routers.BindRoutes)
}
