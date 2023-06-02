package middleware

import (
	"lapson_go_api_sample/config"
	"lapson_go_api_sample/pkg/algoAuth"
	"lapson_go_api_sample/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var middlewareLogger *logrus.Entry = logger.GetLogger("middleware")

func AuthorizeAlgo(config config.APIServerConfig, auth_type string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		match := false
		var err error
		if auth_type == "worker" {
			match, err = algoAuth.ComparePasswordAndHash(config.WorkerPassword, authHeader)
		} else {
			match, err = algoAuth.ComparePasswordAndHash(config.MiniGamePassword, authHeader)
		}
		if !match || err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
