package handler

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetBuilding(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		handlerLogger.Errorf("GetBuilding Error: %v", err)
		handlerLogger.Errorf("id parameter not integer: %v", c.Param("id"))
		h.StatusInternalServerError(c.Writer, c.Request)
	}
	building_obj, error := h.controller.GetBuilding(id)
	
	if error == nil {
		c.IndentedJSON(http.StatusOK, building_obj)
	} else {
		handlerLogger.Errorf("Building Token ID not exist: %v", error)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "Failed",
			"errMsg": "Building Token id not exist",
		})
	}
}

func (h *Handler) GetListBuilding(c *gin.Context) {
	buildings, error := h.controller.GetListBuilding()
	if error == nil {
		c.IndentedJSON(http.StatusOK, buildings)
	} else {
		handlerLogger.Errorf("Get buildings list failed: %v", error)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "Failed",
			"errMsg": "Get buildings list failed",
		})
	}
}

