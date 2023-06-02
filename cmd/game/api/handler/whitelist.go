package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetWhitelistHandler(c *gin.Context) {
	wallet_address := c.Param("wallet")
	upper_wallet := strings.ToUpper(wallet_address)
	is_whitelist, err := h.controller.GetWhitelist(upper_wallet)
	if is_whitelist {
		c.JSON(http.StatusOK, gin.H{
			"status": "succeed",
		})
	} else {
		handlerLogger.Errorf("The wallet address is not in the whitelist: %v", err)
		c.JSON(http.StatusNotFound, gin.H{
			"status": "failed",
			"errMsg": "wallet is not in the whitelist or the wallet is not valid",
		})
	}
}
