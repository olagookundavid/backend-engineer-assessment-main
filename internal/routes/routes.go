package routes

import (
	"expvar"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/masena-dev/bookstore-api/cmd/api"
)

func Routes(app *api.Application) http.Handler {
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(app.NotFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.MethodNotAllowedResponse)

	//Healthcheck
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.HealthcheckHandler)
	//Metrics
	router.Handler(http.MethodGet, "/v1/debug/vars", expvar.Handler())

	//Middleware
	return app.Metrics(app.RecoverPanic(app.RateLimit(router)))
}
