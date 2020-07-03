package controllers

import (
	"net/http"

	"github.com/vanjoechua/orderbackend/api/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome To My Developer Test API")
}
