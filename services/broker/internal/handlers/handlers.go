package handlers

import (
	"net/http"

	"broker/pkg/responses"
	"github.com/gin-gonic/gin"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Broker(c *gin.Context) {
	response := responses.JSONResponse{
		Error:   false,
		Message: "Hit the broker",
	}

	c.JSON(http.StatusOK, response)
}
