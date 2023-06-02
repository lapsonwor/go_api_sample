package main

import (
	"encoding/json"
	"fmt"
	"lapson_go_api_sample/cmd/game/core"
	. "lapson_go_api_sample/common"
	"lapson_go_api_sample/config"
	"lapson_go_api_sample/pkg/logger"
	"os"
	"path"
	"time"

	"github.com/coreos/go-systemd/daemon"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// @version 1.0

// @securityDefinitions.apikey TokenAuth
// @in header
// @name Authorization

const SOFTWARE_NAME = "API Backend"

var (
	VERSION_MAJOR       = "UNKNOWN_VERSION_MAJOR"
	VERSION_MINOR       = "UNKNOWN_VERSION_MINOR"
	VERSION_MAINTENANCE = "UNKNOWN_VERSION_MAINTENANCE"
	VERSION_BUILD       = "UNKNOWN_VERSION_BUILD"
	GIT_BRANCH          = "UNKNOWN_GIT_BRANCH"
	GIT_COMMIT_ID       = "UNKNOWN_GIT_COMMIT_ID"
	BUILD_DATE          = time.Now().Local().Format(time.RFC3339)
)

var cfgFile string
var mainLogger *logrus.Entry = logger.GetLogger("main")

func init() {
	cobra.OnInitialize()
	RootCmd.AddCommand(versionCmd)
	RootCmd.AddCommand(configCmd)
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ./config.json)")
}

var RootCmd = &cobra.Command{
	Use:   SOFTWARE_NAME,
	Short: SOFTWARE_NAME + " is backend api server",
	Long:  "Complete documentation is available at gitlab",
	Run: func(cmd *cobra.Command, args []string) {
		polkaApiServerStart()
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of " + SOFTWARE_NAME,
	Long:  "All software has versions. This is " + SOFTWARE_NAME + `'s`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Software Name       :", SOFTWARE_NAME)
		fmt.Println("Version Major       :", VERSION_MAJOR)
		fmt.Println("Version Minor       :", VERSION_MINOR)
		fmt.Println("Version Maintenance :", VERSION_MAINTENANCE)
		fmt.Println("Version Build       :", VERSION_BUILD)
		fmt.Println("Git Branch          :", GIT_BRANCH)
		fmt.Println("Git Commit ID       :", GIT_COMMIT_ID)
		fmt.Println("Build date          :", BUILD_DATE)
	},
}

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Print the configuration of " + SOFTWARE_NAME,
	Run: func(cmd *cobra.Command, args []string) {
		configuration := loadConfiguration()
		if b, err := json.Marshal(*configuration); IsError(err) {
			fmt.Println(err)
		} else {
			PrintJSON(b)
		}
	},
}

func main() {
	if err := RootCmd.Execute(); IsError(err) {
		fmt.Println(err)
	}
}

func polkaApiServerStart() {
	version := VERSION_MAJOR + "." + VERSION_MINOR + "." + VERSION_MAINTENANCE + "." + VERSION_BUILD
	mainLogger.Infoln("--------------------------------------------------")
	mainLogger.Infoln("Starting application server version:", version)

	apiServer := apiServerSetup()
	daemon.SdNotify(false, "READY=1")
	mainLogger.Infoln("--------------------------------------------------")
	mainLogger.Infoln("Finished application server configuration and setup.")

	apiServer.Start()
	mainLogger.Infoln("--------------------------------------------------")
	mainLogger.Infoln("Start serving application server.")

	go watchdog()
	<-quit
}

var quit = make(chan bool)

func watchdog() {
	interval, err := daemon.SdWatchdogEnabled(false)
	if IsError(err) || interval == 0 {
		mainLogger.Warn("Watchdog is not enabled.")
		return
	}
	for {
		daemon.SdNotify(false, "WATCHDOG=1")
		time.Sleep(interval / 3)
	}
}

func loadConfiguration() *config.Configuration {
	configuration := config.Configuration{}
	if len(cfgFile) > 0 {
		if _, err := os.Stat(cfgFile); os.IsNotExist(err) {
			PanicWhenError(err)
		} else {
			mainLogger.Infoln("Load the application server configuration from", cfgFile)
			if tmpConfig := config.LoadConfiguration(cfgFile); tmpConfig == nil {
				PanicWhenError(err)
			} else {
				configuration = *tmpConfig
			}
		}
	} else {
		workingDir, _ := os.Getwd()
		defaultConfigFile := path.Join(workingDir, config.DAFAULT_CONFIG_FILENAME)
		if _, err := os.Stat(defaultConfigFile); os.IsNotExist(err) {
			PanicWhenError(err)
		} else {
			if tmpConfig := config.LoadConfiguration(defaultConfigFile); tmpConfig == nil {
				PanicWhenError(err)
			} else {
				configuration = *tmpConfig
			}
		}
	}
	return &configuration
}

func apiServerSetup() *core.PolkaApiServer {
	if configuration := loadConfiguration(); configuration != nil {
		return core.NewPolkaApiServer(configuration)
	} else {
		mainLogger.Errorln("The application server configuration is not recognized.")
		os.Exit(2)
	}
	return nil
}
