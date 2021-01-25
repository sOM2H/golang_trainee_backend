package config

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	DSN string `yaml:"dsn"`
}

func New(filePath string) *Config {
	if filePath == "" {
		filePath = "config.yml"

	}
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		log.Println("Config file does not exist:", filePath)
		log.Fatalln(err)
	}
	configFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Println("An error occurred while trying to read the config file:", filePath)
		log.Fatalln(err)
	}

	var config Config

	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		log.Println("Unable to parse contents of YAML config file:", filePath)
		log.Fatalln(err)
	}

	return &config
}
