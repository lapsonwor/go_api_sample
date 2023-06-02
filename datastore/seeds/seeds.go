package main

import (
	"context"
	. "lapson_go_api_sample/common"
	"lapson_go_api_sample/config"
	"lapson_go_api_sample/datastore"
	"lapson_go_api_sample/datastore/sqlc"
	"lapson_go_api_sample/pkg/logger"
	"log"
	"os"
	"path"

	"github.com/sirupsen/logrus"
)

var seedLogger *logrus.Entry = logger.GetLogger("SEED")

func main() {
	config := loadConfiguration()
	db := datastore.NewPostgreSQL(config.DatabaseConfig)
	//Create Rarity
	var rarities = []string{
		"Common", "Rare", "Epic", "Legendary", "Mythical",
	}
	for _, rarity := range rarities {
		_, err := db.Queries.CreateRarity(context.Background(), rarity)
		if err != nil {
			seedLogger.Errorln(err)
		}
	}

	//Create Character
	var characters = []sqlc.CreateCharacterParams{}
	characters = append(characters, sqlc.CreateCharacterParams{
		Name:          "Ameli",
		AttackRange:   1,
		EquipmentSlot: 1,
		CountryID:     1,
		GenderID:      1,
		AttributeID:   1,
	})

	for _, character := range characters {
		_, err := db.Queries.CreateCharacter(context.Background(), character)
		if err != nil {
			seedLogger.Errorln(err)
		}
	}

}

var cfgFile string

func loadConfiguration() *config.Configuration {
	configuration := config.Configuration{}
	if len(cfgFile) > 0 {
		if _, err := os.Stat(cfgFile); os.IsNotExist(err) {
			PanicWhenError(err)
		} else {
			log.Print("Load the application server configuration from", cfgFile)
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
