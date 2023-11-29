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

func ReadConfigYAML() (Config, error) {
	data, err := os.ReadFile("./crm_config.yaml")
	if err != nil {
		return Config{}, err
	}

	// unmarshall config file into Config struct
	var config Config
	err = yaml.Unmarshal([]byte(data), &config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}

func GetDBConnectionString(config Config) (string, error) {
	config, err := ReadConfigYAML()
	if err != nil {
		return "", err
	}

	dsn := fmt.Sprintf("host=%s port=%s dbname=%s sslmode=%s user=%s password=%s \n", config.Db.Pg["dbHost"], config.Db.Pg["dbPort"],
		config.Db.Pg["dbName"], config.Db.Pg["dbSslMode"], config.Db.Pg["dbUser"], config.Db.Pg["dbPassword"])

	return dsn, nil
}
