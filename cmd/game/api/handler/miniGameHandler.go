package handler

import (
	"encoding/base64"
	"encoding/json"
	"lapson_go_api_sample/models"
	"lapson_go_api_sample/pkg/goEth"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *Handler) SaveMiniGameMark(c *gin.Context) {
	//Get publicAddr, encryptedMessage and signature and verify wallet address
	publicAddr := c.PostForm("publicAddr")
	signatureHash := c.PostForm("signatureHash")
	encryptedMessage := c.PostForm("encryptedMessage")
	verified := goEth.VerifySign(publicAddr, signatureHash, encryptedMessage)

	if !verified {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "failed",
			"errMsg": "Save MiniGame mark failed. Wallet not verified",
		})
		return
	}

	//decrypt message
	decodeEncryptedByte, err := base64.StdEncoding.DecodeString(encryptedMessage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "failed",
			"errMsg": "Save MiniGame mark failed. Failed to decode message",
		})
		return
	}
	rsaPrivateKey := goEth.PemToPrivateKey(goEth.PrivateKey)
	messageByte := goEth.DecryptWithPrivateKey(decodeEncryptedByte, rsaPrivateKey)
	plaintext := string(messageByte[:])
	handlerLogger.Info(plaintext)

	//handle mark data
	miniGameJson := models.MiniGamePostJson{}
	json.Unmarshal(messageByte, &miniGameJson)
	wallet_address := miniGameJson.Wallet
	mark := miniGameJson.Mark
	handlerLogger.Info(miniGameJson)
	if wallet_address != "" {
		upper_wallet := strings.ToUpper(wallet_address)
		mark_record, err := h.controller.SaveMiniGameMark(upper_wallet, mark)
		if err == nil {
			c.Data(http.StatusOK, "application/json", []byte(mark_record))
		} else {
			handlerLogger.Errorf("Save MiniGame Mark failed: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "failed",
				"errMsg": "Save MiniGame mark failed." + err.Error(),
			})
		}
	} else {
		handlerLogger.Errorf("Wallet address is empty")
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "failed",
			"errMsg": "Save MiniGame mark failed. Wallet address cannot be empty.",
		})
	}
}

func (h *Handler) GetMiniGameMarks(c *gin.Context) {
	mark_records, err := h.controller.GetMiniGameMarks()
	if err == nil {
		c.Data(http.StatusOK, "application/json", []byte(mark_records))
	} else {
		handlerLogger.Errorf("Get MiniGame Mark failed: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "failed",
			"errMsg": "Get mark failed. Please contact admin for futher assistance.",
		})
	}
}

func (h *Handler) GetMiniGameRankByWallet(c *gin.Context) {
	wallet_json := models.MiniGameWalletJson{}
	c.BindJSON(&wallet_json)
	wallet_address := strings.ToUpper(wallet_json.Wallet)
	rank, err := h.controller.GetMiniGameRankByWallet(wallet_address)
	if err == nil {
		c.Data(http.StatusOK, "application/json", []byte(rank))
	} else {
		handlerLogger.Errorf("Get MiniGame Mark failed: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "failed",
			"errMsg": "Get Rank failed. Please contact admin for futher assistance.",
		})
	}
}

func (h *Handler) GetMiniGameTop10(c *gin.Context) {
	top10scores, err := h.controller.GetMiniGameTop10()
	if err == nil {
		c.Data(http.StatusOK, "application/json", []byte(top10scores))
	} else {
		handlerLogger.Errorf("Get Top 10 scores failed: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "failed",
			"errMsg": "Get mark failed. Please contact admin for futher assistance.",
		})
	}
}
