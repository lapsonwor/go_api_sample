package core

import (
	"lapson_go_api_sample/cmd/game/api"
	//. "lapson_go_api_sample/common"
	"lapson_go_api_sample/cmd/game/api/handler"
	"lapson_go_api_sample/cmd/game/controller"
	"lapson_go_api_sample/config"
	"lapson_go_api_sample/pkg/logger"

	"github.com/sirupsen/logrus"
)

var coreLogger *logrus.Entry = logger.GetLogger("core")

type PolkaApiServer struct {
	config     *config.Configuration
	controller *controller.Controller
	handler    *handler.Handler
	api        *api.APIServer
}

func NewPolkaApiServer(configuration *config.Configuration) *PolkaApiServer {

	// coreLogger.Infoln("--------------------------------------------------")
	// coreLogger.Infoln("Create new fetcher.")
	// coreFetcher := fetcher.New(configuration.TargetServer)

	coreLogger.Infoln("--------------------------------------------------")
	coreLogger.Infoln("Create new controller.")
	coreController := controller.New(configuration)

	coreLogger.Infoln("--------------------------------------------------")
	coreLogger.Infoln("Create new handler.")
	coreHandler := handler.New(coreController)

	coreLogger.Infoln("--------------------------------------------------")
	coreLogger.Infoln("Create new api server.")
	coreAPIServer := api.New(configuration.APIServer, coreHandler)

	return &PolkaApiServer{
		config: configuration,
		// fetcher:    coreFetcher,
		controller: coreController,
		handler:    coreHandler,
		api:        coreAPIServer,
	}
}

func (server *PolkaApiServer) Start() {
	// response, err := server.fetcher.Fetch("/map/list-item/5")
	// if err != nil {
	// 	coreLogger.Infoln(err)
	// } else {
	// 	PrintJSON(response)
	// }
	server.api.Start()
}
