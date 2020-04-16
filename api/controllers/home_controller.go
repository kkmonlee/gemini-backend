package controllers

import (
	"net/http"

	"github.com/kkmonlee/gemini-backend/gemini-backend/api/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome to Gemini API")
}
