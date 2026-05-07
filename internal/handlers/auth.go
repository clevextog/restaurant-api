package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/clevextog/restaurant-api/internal/user"
)

type AuthHandler struct {
	service *user.AuthService
}

func NewAuthHandler(service *user.AuthService) *AuthHandler {
	return &AuthHandler{
		service: service,
	}
}

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (s *AuthHandler) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var req RegisterRequest
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&req)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, ErrorResponse{Error: "bad request"})
		return
	}
	if req.Email == "" || req.Password == "" {
		writeJSON(w, http.StatusBadRequest, ErrorResponse{Error: "bad request"})
		return
	}
	err = s.service.Register(req.Email, req.Password)
	if errors.Is(err, user.ErrUserExists) {
		writeJSON(w, http.StatusConflict, ErrorResponse{Error: "invalid credentials"})
		return
	}
	writeJSON(w, http.StatusCreated, MessageResponse{Message: "successfull register"})
}
