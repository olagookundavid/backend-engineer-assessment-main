package routes

import (
	"expvar"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"

	"github.com/masena-dev/bookstore-api/cmd/api"
)

func Routes(app *api.Application) http.Handler {
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(app.NotFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.MethodNotAllowedResponse)

	//Handlers
	authorHandler := app.Handlers.AuthorsHandler
	bookHandler := app.Handlers.BooksHandler

	//Healthcheck
	router.HandlerFunc(http.MethodGet, "/api/v1/healthcheck", app.HealthcheckHandler)

	//Metrics
	router.Handler(http.MethodGet, "/api/v1/debug/vars", expvar.Handler())

	//Authors
	router.HandlerFunc(http.MethodGet, "/api/v1/authors/:id", authorHandler.GetAuthor)
	router.HandlerFunc(http.MethodGet, "/api/v1/authors", authorHandler.GetAllAuthors)
	router.HandlerFunc(http.MethodGet, "/api/v1/authors/:id/stats", authorHandler.GetAuthorStats)

	//Books
	router.HandlerFunc(http.MethodGet, "/api/v1/books/:id", bookHandler.GetBook)
	router.HandlerFunc(http.MethodGet, "/api/v1/books", bookHandler.GetAllBooks)
	router.HandlerFunc(http.MethodPut, "/api/v1/books/:id", bookHandler.UpdateBook)
	router.HandlerFunc(http.MethodPost, "/api/v1/books", bookHandler.CreateBook)
	router.HandlerFunc(http.MethodDelete, "/api/v1/books/:id", bookHandler.DeleteBook)

	//Middleware
	// return app.Metrics(app.RecoverPanic(app.RateLimit(router)))
	MiddlewareChain := alice.New(app.Metrics, app.RecoverPanic, app.RateLimit)
	// auth := MiddlewareChain.Append(app.Authentication)

	return MiddlewareChain.Then(router)
}
