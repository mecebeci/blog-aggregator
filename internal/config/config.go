package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	DBUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (Config, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return Config{}, err
	}

	configPath := filepath.Join(home, ".gatorconfig.json")

	data, err := os.ReadFile(configPath)
	if err != nil {
		return Config{}, err
	}

	var cfg Config
	err = json.Unmarshal(data, &cfg)
	if err != nil {
		return Config{}, err
	}
	return cfg, nil
}

func (c *Config) SetUser(username string) error {
	c.CurrentUserName = username

	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	configPath := filepath.Join(home, ".gatorconfig.json")

	data, err := json.Marshal(c)
	if err != nil {
		return err
	}

	err = os.WriteFile(configPath, data, 0644)
	if err != nil {
		return err
	}
	return nil
}
