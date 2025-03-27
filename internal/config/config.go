package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DBUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (Config, error) {
	filepath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	// Read the file content
	file, err := os.Open(filepath)
	if err != nil {
		return Config{}, err
	}
	defer file.Close()

	// Decode the file content into a Config struct
	c := Config{}
	err = json.NewDecoder(file).Decode(&c)
	if err != nil {
		return Config{}, err
	}

	return c, nil
}

func (c *Config) SetUser(username string) error {
	c.CurrentUserName = username
	return write(*c)
}

func getConfigFilePath() (string, error) {
	// Get the home directory of the user
	projectRoot, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	// Join the project root with the config file name
	return filepath.Join(projectRoot, configFileName), nil
}

func write(c Config) error {
	filepath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	// Write the Config struct to the file
	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	err = json.NewEncoder(file).Encode(c)
	if err != nil {
		return err
	}
	return nil
}
