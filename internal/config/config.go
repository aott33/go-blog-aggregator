package config

import (
	"encoding/json"
	"path/filepath"
	"io"
	"os"
)

func Read() (Config, error){
	var cfg Config
	
	fullPath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	jsonFile, err := os.Open(fullPath)	
	if err != nil {
		return Config{}, err
	}

	defer jsonFile.Close()

	byteVal, err := io.ReadAll(jsonFile)
	if err != nil {
		return Config{}, err
	}

	err = json.Unmarshal(byteVal, &cfg)
	if err != nil {
		return Config{}, err
	}

	return cfg, nil
}

func (c *Config) SetUser(userName string) error {
	c.CurrentUserName =  userName
	err := write(*c)
	if err != nil {
		return err
	}

	return nil
}

func getConfigFilePath() (string, error) {

	homeDir, err := os.UserHomeDir()

	if err != nil {
		return "", err
	}

	fullPath := filepath.Join(homeDir, configFileName)	

	return fullPath, nil
}

func write(cfg Config) error {
	fullPath, err := getConfigFilePath()
	if err != nil {
		return err 
	}

	jsonData, err := json.Marshal(cfg)
	if err != nil {
		return err
	}
	
	jsonFile, err := os.Create(fullPath)	
	if err != nil {
		return err
	}

	defer jsonFile.Close()

	_, err = jsonFile.Write(jsonData)
	if err != nil {
		return err
	}

	return nil
}
