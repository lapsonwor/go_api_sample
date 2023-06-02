package handler

import (
	"lapson_go_api_sample/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetCardNFT(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "Failed",
			"errMsg": "Token id not valid",
		})
		return
	}
	character, error := h.controller.GetCardNFT(id)

	if error == nil {
		c.IndentedJSON(http.StatusOK, character)
	} else {
		handlerLogger.Errorf("Token ID not exist: %v", error)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "Failed",
			"errMsg": "Token id not exist",
		})
	}
}

func (h *Handler) CreateCardFromWorker(c *gin.Context) {
	workerjson := models.WorkerPostJson{}
	c.BindJSON(&workerjson)

	message, error := h.controller.CreateCardFromWorker(workerjson)
	if error == nil {
		c.Data(http.StatusOK, "application/json", []byte(message))
	} else {
		handlerLogger.Infof(error.Error())
		c.Data(http.StatusBadRequest, "application/json", []byte(message))
	}
}
