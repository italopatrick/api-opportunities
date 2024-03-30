package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/italopatrick/api-opportunities/schemas"
)

func ShowOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id:", "query parameter").Error())
		return
	}
	opening := schemas.Opening{}
	if err := db.First(&opening).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "opening not found")
		return
	}

	sendSuccess(ctx, "show opening", opening)

}
