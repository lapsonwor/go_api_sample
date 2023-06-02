package api

import (
	"lapson_go_api_sample/cmd/game/api/handler"
	"lapson_go_api_sample/cmd/game/api/middleware"
	"lapson_go_api_sample/config"
	"lapson_go_api_sample/pkg/logger"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var apiLogger *logrus.Entry = logger.GetLogger("api")

type APIServer struct {
	config  config.APIServerConfig
	handler *handler.Handler
}

func New(apiServerConfig config.APIServerConfig, handler *handler.Handler) *APIServer {
	return &APIServer{
		config:  apiServerConfig,
		handler: handler,
	}
}

func corsConfig() cors.Config {
	corsConf := cors.DefaultConfig()
	corsConf.AllowAllOrigins = true
	corsConf.AllowMethods = []string{"GET", "POST", "DELETE", "OPTIONS", "PUT"}
	corsConf.AllowHeaders = []string{"Authorization", "Content-Type", "Upgrade", "Origin",
		"Connection", "Accept-Encoding", "Accept-Language", "Host", "Access-Control-Request-Method", "Access-Control-Request-Headers"}
	return corsConf
}

func (api *APIServer) Start() {
	apiLogger.Info("Start loading route map.")
	gin.SetMode(gin.DebugMode)
	router := gin.Default()

	// testing only
	router.Use(cors.New(corsConfig()))
	router.GET("/", api.handler.Hello)
	router.GET("/cardNFT/:id", api.handler.GetCardNFT)
	router.GET("/whitelist", api.handler.GetWhitelistHandler)
	router.GET("/whitelist/:wallet", api.handler.GetWhitelistHandler)
	router.GET("/buildings/", api.handler.GetListBuilding)
	router.GET("/building/:id", api.handler.GetBuilding)

	workerAuth := router.Group("/cardNFT/")
	workerAuth.Use(middleware.AuthorizeAlgo(api.config, "worker"))
	{
		workerAuth.POST("/createCardFromWorker", api.handler.CreateCardFromWorker)
	}
	miniGameAuth := router.Group("/mini-game/")
	miniGameAuth.Use(middleware.AuthorizeAlgo(api.config, "miniGame"))
	{
		miniGameAuth.POST("/saveMark", api.handler.SaveMiniGameMark)
		miniGameAuth.GET("/marks", api.handler.GetMiniGameMarks)
		miniGameAuth.POST("/getRankByWallet", api.handler.GetMiniGameRankByWallet)
		miniGameAuth.GET("/getTop10Scores", api.handler.GetMiniGameTop10)
	}
	router.Run("0.0.0.0:" + strconv.Itoa(api.config.Port))
}
