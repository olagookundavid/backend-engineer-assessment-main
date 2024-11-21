/*
mkdir -p bin cmd/api internal migrations remote
cmd/api/main.go

cd internal/sql/migrations/
goose postgres postgres://djjsagev:WG11sRXwe2q1C0I9-3XhTZywTnhbZQPJ@stampy.db.elephantsql.com/djjsagev up
goose postgres postgres://itojudb:itojudb@localhost/itojudb up
*/
package main

import (
	"os"
	"sync"

	_ "github.com/lib/pq"
	"github.com/masena-dev/bookstore-api/cmd/api"
	"github.com/masena-dev/bookstore-api/internal/jsonlog"

	// "github.com/masena-dev/bookstore-api/internal/models"
	"github.com/masena-dev/bookstore-api/internal/server"
)

func main() {
	dbUrl := loadDbUrl()
	cfg := flagSetup(dbUrl)
	displayVersion("version")

	logger := jsonlog.New(os.Stdout, jsonlog.LevelInfo)
	db, err := openDB(*cfg)
	if err != nil {
		logger.PrintFatal(err, nil)
	}
	logger.PrintInfo("database connection pool established", nil)

	defer db.Close()

	expvarSetup(db)

	app := &api.Application{
		Wg:     sync.WaitGroup{},
		Config: *cfg,
		Logger: logger,
		//TODO
		// Models: models.NewModels(db),
	}

	err = server.Serve(app)
	if err != nil {
		logger.PrintFatal(err, nil)

	}

}
