package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func (cfg Config) Repr() {
	fmt.Printf("Config{\n	DbURL: %s,\n	CurrentUserName: %s\n}\n", cfg.DbURL, cfg.CurrentUserName)
}

func (cfg Config) SetUser(username string) error {
	if username == "" {
		return errors.New("no username provided")
	}

	cfg.CurrentUserName = username
	err := write(cfg)
	if err != nil {
		return err
	}

	return nil

}

func write(cfg Config) error {
	path, err := getConfigFilePath()
	if err != nil {
		return err
	}
	marshaledData, err := json.Marshal(cfg)
	if err != nil {
		return err
	}
	err = os.WriteFile(path, marshaledData, 0666)
	if err != nil {
		return err
	}
	return nil
}

func getConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	if home == "" {
		return "", errors.New("no home directory found")
	}

	return filepath.Join(home, configFileName), nil
}

func Read() (Config, error) {
	var cfg Config
	path, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	readResult, err := os.ReadFile(path)
	if err != nil {
		return Config{}, err
	}

	err = json.Unmarshal(readResult, &cfg)
	if err != nil {
		return Config{}, err
	}
	return cfg, nil
}
