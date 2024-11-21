package api

import (
	"sync"

	"github.com/masena-dev/bookstore-api/internal/handlers"
	"github.com/masena-dev/bookstore-api/internal/jsonlog"
	// "github.com/masena-dev/bookstore-api/internal/services"
)

type Application struct {
	Handlers handlers.Handlers
	Config   Config
	Logger   *jsonlog.Logger
	Wg       sync.WaitGroup
}

type Config struct {
	Port int
	Env  string
	Db   struct {
		Dsn          string
		MaxOpenConns int
		MaxIdleConns int
		MaxIdleTime  string
	}
	Limiter struct {
		Rps     float64
		Burst   int
		Enabled bool
	}
	Cors struct {
		TrustedOrigins []string
	}
}
