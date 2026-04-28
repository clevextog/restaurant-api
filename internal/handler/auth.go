package handler

import "github.com/clevextog/restaurant-api/internal/user"

type AuthHandler struct {
	service *user.AuthService
}

func NewAuthHandler(service *user.AuthService) *AuthHandler {
	return &AuthHandler{
		service: service,
	}
}
