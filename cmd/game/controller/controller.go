package controller

import (
	"encoding/json"
	"lapson_go_api_sample/config"
	"lapson_go_api_sample/datastore"
	"lapson_go_api_sample/metadata"
	"lapson_go_api_sample/models"
	"os"
	"path/filepath"

	"reflect"
	"regexp"
	"strings"

	_ "lapson_go_api_sample/datastore"
	"lapson_go_api_sample/pkg/logger"

	"github.com/sirupsen/logrus"
)

var controllerLogger *logrus.Entry = logger.GetLogger("controller")

type Controller struct {
	db        *datastore.Datastore
	allconfig *config.Configuration
}

func New(allconfig *config.Configuration) *Controller {
	db := datastore.NewPostgreSQL(allconfig.DatabaseConfig)
	return &Controller{
		db:        db,
		allconfig: allconfig,
	}
}

func responseErrMsg(errorScope string, err error) (string, error) {
	errMsg := err.Error()
	if strings.HasPrefix(errMsg, "pq: duplicate key") {
		errMsg = errorScope + " Record already exists"
	}
	message := models.ResponseErrJson{
		Status: "failed",
		ErrMsg: errorScope + ": " + errMsg,
	}
	errjson, errorMar := json.MarshalIndent(message, "", "\t")
	if errorMar != nil {
		return errMsg, errorMar
	}
	return string(errjson), err
}

func generateJsonFile(card metadata.Metadata, filename_json string) {
	dir := "./ipfsJson/"
	dst, err := os.Create(filepath.Join(dir, filepath.Base(filename_json)))
	if err != nil {
		controllerLogger.Infoln("Unable to create json file")
	}
	defer dst.Close()
	jsonStr, err := json.MarshalIndent(card, "", "\t")
	if err != nil {
		controllerLogger.Infoln("[GenerateJsonFile] Json failed to Marshal")
	}
	_, err = dst.Write(jsonStr)
	if err != nil {
		controllerLogger.Infoln("[GenerateJsonFile] Failed to write json file")
	}
}

func getCardField(v interface{}, field string) string {
	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(field)
	return f.String()
}

func getFileName(str string) string {
	// convert every letter to lower case
	newStr := str

	// convert all spaces/tab to underscore
	regExp := regexp.MustCompile("[[:space:][:blank:]]")
	newStrByte := regExp.ReplaceAll([]byte(newStr), []byte("_"))
	// and remove underscore from beginning and ending

	newStr = strings.TrimPrefix(string(newStrByte), "_")
	newStr = strings.TrimSuffix(newStr, "_")

	return newStr
}
