package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) getBars(c *gin.Context) {
	fmt.Println("getBars")

	bars, err := h.services.BarComm.GetBars()
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, bars)
}

func (h *Handler) getBarName(id_place int) (string, error) {
	fmt.Println("getBarName")

	barName, err := h.services.BarComm.GetBarName(id_place)
	if err != nil {
		return "", nil
	}
	return barName, nil
}
