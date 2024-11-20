package api

import (
	"encoding/json"
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/masena-dev/bookstore-api/internal/db"
	"log/slog"
	"net/http"
	"strconv"
)

type Server struct {
	queries *db.Queries
	logger  *slog.Logger
}

func NewServer(logger *slog.Logger, db *db.Queries) *Server {
	return &Server{
		queries: db,
		logger:  logger,
	}
}

func (s *Server) NewRouter() http.Handler {
	mux := http.NewServeMux()

	// Authors endpoints
	mux.HandleFunc("GET /api/v1/authors", s.handleListAuthors)
	mux.HandleFunc("GET /api/v1/authors/{id}/stats", s.handleGetAuthorStats)

	// Books endpoints
	mux.HandleFunc("GET /api/v1/books", s.handleListBooks)
	mux.HandleFunc("POST /api/v1/books", s.handleCreateBook)
	mux.HandleFunc("GET /api/v1/books/{id}", s.handleGetBook)
	mux.HandleFunc("PUT /api/v1/books/{id}", s.handleUpdateBook)
	mux.HandleFunc("DELETE /api/v1/books/{id}", s.handleDeleteBook)

	return mux
}

func (s *Server) respond(w http.ResponseWriter, status int, data interface{}) {
	if data == nil {
		w.WriteHeader(status)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		s.logger.Error("failed to encode response", "error", err)
	}
}

func (s *Server) respondError(w http.ResponseWriter, status int, message string) {
	s.logger.Error("error", "error_message", message)
	errorResponse := struct {
		Message string `json:"message"`
	}{
		Message: message,
	}
	s.respond(w, status, errorResponse)
}

// Authors handlers
func (s *Server) handleListAuthors(w http.ResponseWriter, r *http.Request) {
	s.respond(w, http.StatusNoContent, nil)
}

func (s *Server) handleGetAuthorStats(w http.ResponseWriter, r *http.Request) {
	_, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		s.respondError(w, http.StatusBadRequest, "Invalid author ID")
		return
	}
	s.respond(w, http.StatusNoContent, nil)
}

// Books handlers
func (s *Server) handleListBooks(w http.ResponseWriter, r *http.Request) {
	s.respond(w, http.StatusNoContent, nil)
}

func (s *Server) handleCreateBook(w http.ResponseWriter, r *http.Request) {
	var req CreateBookRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		s.respondError(w, http.StatusBadRequest, "Invalid request format")
		return
	}

	if err := req.Validate(); err != nil {
		s.respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	s.respond(w, http.StatusNoContent, nil)
}

func (s *Server) handleGetBook(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		s.respondError(w, http.StatusBadRequest, "Invalid book ID")
		return
	}

	book, err := s.queries.GetBook(r.Context(), id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			s.respondError(w, http.StatusNotFound, "book not found")
			return
		}
		s.respondError(w, http.StatusInternalServerError, err.Error())
	}

	s.respond(w, http.StatusOK, book)
}

func (s *Server) handleUpdateBook(w http.ResponseWriter, r *http.Request) {
	_, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		s.respondError(w, http.StatusBadRequest, "Invalid book ID")
		return
	}
	s.respond(w, http.StatusNoContent, nil)
}

func (s *Server) handleDeleteBook(w http.ResponseWriter, r *http.Request) {
	_, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		s.respondError(w, http.StatusBadRequest, "Invalid book ID")
		return
	}
	s.respond(w, http.StatusNoContent, nil)
}
