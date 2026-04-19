package main

import (
	"encoding/json"
	"net/http"
)

func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func fallback(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusNotFound, ErrorResponse{Error: "not found"})
}

func RegistrationHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var req RegistrationRequest
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	err := dec.Decode(&req)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, ErrorResponse{Error: "bad request"})
		return
	}
	err = Register(req.Email, req.Password)
	if err == ErrUserExists {
		writeJSON(w, http.StatusConflict, ErrorResponse{Error: "user already exists"})
		return
	}
	if err == ErrEmailValidation {
		writeJSON(w, http.StatusBadRequest, ErrorResponse{Error: "email validation failed"})
		return
	}
	if err == ErrPasswordValidation {
		writeJSON(w, http.StatusBadRequest, ErrorResponse{Error: "password validation failed"})
		return
	}
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, ErrorResponse{Error: "internal server error"})
		return
	}
	writeJSON(w, http.StatusCreated, MessageResponse{Message: "registration was successful"})
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var req LoginRequest
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	err := dec.Decode(&req)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, ErrorResponse{Error: "bad request"})
		return
	}
	err = Login(req.Email, req.Password)
	if err == ErrInvalidCredentials {
		writeJSON(w, http.StatusUnauthorized, ErrorResponse{Error: ErrInvalidCredentials.Error()})
		return
	}
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, ErrorResponse{Error: "internal server error"})
		return
	}
	writeJSON(w, http.StatusOK, MessageResponse{Message: "login was successful"})
}
