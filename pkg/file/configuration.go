package file

import (
	"fmt"
	"github.com/tkanos/gonfig"
	"os"
)

type Config struct {
	LogDirectory string `json:"log_directory"`
	LogFilename  string `json:"log_filename"`
	Database     struct {
		Charset  string `json:"charset"`
		Host     string `json:"host"`
		Username string `json:"username"`
		Password string `json:"password"`
		Schema   string `json:"schema"`
	}
}

func LoadConfiguration(configFilePath string) Config {
	conf := Config{}
	err := gonfig.GetConf(configFilePath, &conf)

	if err != nil {
		fmt.Println("Failed parsing config file " + configFilePath)
		os.Exit(1)
	}

	fmt.Println("Initialized configuration")
	return conf
}
