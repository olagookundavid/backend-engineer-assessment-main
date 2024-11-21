package transport

import (
	"encoding/json"
	"net/http"

	"github.com/masena-dev/bookstore-api/internal/domain"
	"github.com/masena-dev/bookstore-api/internal/helpers"
	"github.com/masena-dev/bookstore-api/internal/services"
	// "github.com/gorilla/mux"
)

type AuthorHandler struct {
	AuthorService *services.AuthorService
}

func NewAuthorHandler(service *services.AuthorService) *AuthorHandler {
	return &AuthorHandler{AuthorService: service}
}

func (h *AuthorHandler) GetAuthor(w http.ResponseWriter, r *http.Request) {
	id, err := helpers.ReadIDParam(r)

	user, err := h.AuthorService.GetAuthor(r.Context(), id)
	if err != nil {
		http.Error(w, "Author not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func (h *AuthorHandler) CreateAuthor(w http.ResponseWriter, r *http.Request) {
	var user domain.Author
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	newAuthor, err := h.AuthorService.CreateAuthor(r.Context(), &user)
	if err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(newAuthor)
}
