package application

import (
	"encoding/json"
	"os"

	"github.com/snathale/challenge-hash/calculator/infrastructure"
	"github.com/snathale/challenge-hash/calculator/interface/server"
)

type Config struct {
	Db       infrastructure.Config `mapstructure:"database"`
	Server   server.Config         `mapstructure:"grpc_server"`
	LogLevel int                   `mapstructure:"logLevel"`
}

func NewConfigFile(filename string) error {
	var config Config
	err := generateConfigFile(filename, config)
	if err != nil {
		return err
	}
	return nil
}

func generateConfigFile(filename string, config Config) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	data, err := json.Marshal(config)
	if err != nil {
		return err
	}
	_, err = file.Write(data)
	if err != nil {
		return err
	}
	return nil
}
