package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/italopatrick/api-opportunities/schemas"
)

func ListOpeningsHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing openings")
		return // Return after encountering an error
	}

	// Send response with openings
	sendSuccess(ctx, "list-Openings", openings)
}
