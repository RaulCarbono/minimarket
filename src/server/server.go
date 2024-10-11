package server

import (
	"context"
	"errors"
	"log"

	"github.com/go/mini_market/src/database"
	"github.com/go/mini_market/src/repository"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Config struct {
	Port        string
	DatabaseUrl string
	JWTSecret   string
}

type Server interface {
	Config() *Config
}

type Broker struct {
	config *Config
	router *echo.Router
}

func (b *Broker) Config() *Config {
	return b.config
}

func NewServer(ctx context.Context, config *Config) (*Broker, error) {
	if config.Port == "" {
		return nil, errors.New("port is required")
	}
	if config.DatabaseUrl == "" {
		return nil, errors.New("DatabaseUrl is required")
	}

	if config.JWTSecret == "" {
		return nil, errors.New("JWTSecret is required")
	}

	broker := &Broker{
		config: config,
		router: echo.New().Router(),
	}

	return broker, nil
}

func (b *Broker) Start(binder func(s Server, r *echo.Router)) {
	e := echo.New()
	b.router = e.Router()
	binder(b, b.router)

	e.Use(middleware.CORS())
	repo, err := database.DBConnection(b.config.DatabaseUrl)
	if err != nil {
		log.Fatal(err)
	}
	repository.SetRepository(repo)

	log.Println("Starting server on port", b.Config().Port)

	e.Logger.Fatal(e.Start(b.config.Port))
}
