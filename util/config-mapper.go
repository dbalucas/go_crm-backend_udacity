package util

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Db struct {
		Pg map[string]string `yaml:"postgres"`
	}
	Server struct {
		Http map[string]string `yaml:"http"`
	}
}

func ReadConfigYAML() (string, error) {
	data, err := os.ReadFile("./server_config.yaml")
	if err != nil {
		return "", err
	}

	// unmarshall config file into struct
	var config Config
	err = yaml.Unmarshal([]byte(data), &config)
	if err != nil {
		return "", err
	}

	for k, v := range config.Db.Pg {
		fmt.Printf("key %s: %s\n", k, v)
	}

	return fmt.Sprintf("host=%s port=%s dbname=%s sslmode=%s user=%s \n", config.Db.Pg["dbHost"], config.Db.Pg["dbPort"],
		config.Db.Pg["dbName"], config.Db.Pg["dbSslMode"], config.Db.Pg["dbUser"]), nil
}
