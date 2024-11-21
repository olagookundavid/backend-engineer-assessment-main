/*
mkdir -p bin cmd/api internal migrations remote
cmd/api/main.go

cd internal/sql/migrations/
goose postgres postgres://djjsagev:WG11sRXwe2q1C0I9-3XhTZywTnhbZQPJ@stampy.db.elephantsql.com/djjsagev up
goose postgres postgres://itojudb:itojudb@localhost/itojudb up
*/
package main

import (
	"context"
	"os"
	"sync"

	_ "github.com/lib/pq"
	"github.com/masena-dev/bookstore-api/cmd/api"
	"github.com/masena-dev/bookstore-api/internal/db"
	"github.com/masena-dev/bookstore-api/internal/handlers"
	"github.com/masena-dev/bookstore-api/internal/jsonlog"
	"github.com/masena-dev/bookstore-api/internal/server"
)

func main() {
	dbUrl := loadDbUrl()
	cfg := flagSetup(dbUrl)
	displayVersion("version")

	logger := jsonlog.New(os.Stdout, jsonlog.LevelInfo)

	ctx := context.Background()

	pool, err := openDB(*cfg, ctx)
	if err != nil {
		logger.PrintFatal(err, nil)
	}
	defer pool.Close()

	logger.PrintInfo("database connection pool established", nil)

	expvarSetup()

	app := &api.Application{
		Wg:       sync.WaitGroup{},
		Config:   *cfg,
		Logger:   logger,
		Handlers: handlers.NewHandlers(db.New(pool)),
	}

	err = server.Serve(app)
	if err != nil {
		logger.PrintFatal(err, nil)

	}

}
