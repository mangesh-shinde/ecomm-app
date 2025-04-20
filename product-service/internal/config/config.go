package config

import (
	"flag"
	"io"
	"log"
	"os"

	"github.com/mangesh-shinde/ecomm-app/product-service/pkg/models"
	"gopkg.in/yaml.v3"
)

func MustLoad() *models.Config {
	// check if config file path is specified through CONFIG_PATH environment variable
	// if yes, read config file from that path
	// else check if -config argument is specified in command line and read from it
	// if both options are not set, terminate the process
	var configPath string
	configPath = os.Getenv("CONFIG_PATH")

	if configPath == "" {
		confPath := flag.String("config", "", "configuration file for product service")
		flag.Parse()
		configPath = *confPath

		if configPath == "" {
			log.Fatalf("<*> config file path is not specified. Exiting...\n")
		}
	}

	// check if file exists at specified location
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("<*> config file doesn not exist: %s\nExiting...\n", configPath)
	}

	var cfg models.Config
	yamlFile, err := os.Open(configPath)
	if err != nil {
		log.Fatalf("<*> Could not open yaml file: %s\n", configPath)
	}
	defer yamlFile.Close()

	yamlBytes, err := io.ReadAll(yamlFile)
	if err != nil {
		if err != io.EOF {
			log.Fatalf("<*>Unable to read yaml file: %s\n", configPath)
		}
	}

	if err := yaml.Unmarshal(yamlBytes, &cfg); err != nil {
		log.Fatalf("<*> Unable to load config file\n")
	}

	return &cfg

}
