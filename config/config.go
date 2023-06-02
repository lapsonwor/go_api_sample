package config

import (
	"encoding/json"
	"io/ioutil"
	"os"

	. "lapson_go_api_sample/common"
)

const DAFAULT_CONFIG_FILENAME string = "config.json"

type Configuration struct {
	APIServer       APIServerConfig `json:"api_server"`
	DatabaseConfig  DatabaseConfig  `json:"database"`
	MetaJasonConfig MetaJsonConfig  `json:"meta_json"`
}
type DatabaseConfig struct {
	Host        string `json:"host"`
	Port        string `json:"port"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	DbName      string `json:"dbname"`
	SSLMode     string `json:"SSLMode"`
	MaxIdleConn int    `json:"maxIdleConn"`
	MaxOpenConn int    `json:"maxOpenConn"`
}

type APIServerConfig struct {
	Port             int              `json:"port"`
	WorkerPassword   string           `json:"worker_password"`
	MiniGamePassword string           `json:"minigame_password"`
	CertFile         string           `json:"cert_file"`
	KeyFile          string           `json:"key_file"`
	FileServer       FileServerConfig `json:"file_server"`
	SwaggerEnable    bool             `json:"swagger_enable"`
}

type FileServerConfig struct {
	Scheme string `json:"scheme"`
	Host   string `json:"host"`
}

type MetaJsonConfig struct {
	IpfsImgURL           string `json:"ipfs_img_url"`
	BuildingIpfsImageURL string `json:"buildng_ipfs_img_url"`
	ExternalURL          string `json:"external_url"`
}

func LoadConfiguration(name string) *Configuration {
	config := Configuration{}
	file, err := os.Open(name)
	defer file.Close()
	PanicWhenError(err)
	b, err := ioutil.ReadAll(file)
	PrintJSON(b)
	PanicWhenError(err)
	err = json.Unmarshal(b, &config)
	return &config
}
