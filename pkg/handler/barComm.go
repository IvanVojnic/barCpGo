package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) getBars(c *gin.Context) {
	fmt.Println("getBars")

	message, err := h.services.BarComm.GetBars()
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, message)
}
