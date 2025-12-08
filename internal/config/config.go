package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

// State struct- represent "source of truth" state for db config
type State struct {
	CfgPointer *Config
}

// Database configuration info struct representation
type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func (cfg Config) Repr() {
	fmt.Printf("Config{\n	DbURL: %s,\n	CurrentUserName: %s\n}\n", cfg.DbURL, cfg.CurrentUserName)
}

// Set username based on argument then update the user info in the config file
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

// Helper - write config struct to file as JSON object
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

// Helper - fetch config filepath
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

// Return user data from config file as Config struct
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
