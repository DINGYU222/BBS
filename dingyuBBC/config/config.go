package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var (
	config *Config
)

func init() {
	bytes, err := ioutil.ReadFile("./config/config.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = json.Unmarshal(bytes, config)
	if err != nil {
		fmt.Println(err)
		return
	}
}

type Config struct {
	MySQLUserName string `json:"mysqlUsername"`
	MySQLPassword string `json:"mysqlPassword"`
	Database      string `json:"database"`
	MySQLIP       string `json:"ip"`
}

func GetConfig() *Config {
	return config
}
