package config

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"io"
	"os"
)

func Read(file string) Config {
	var cfg Config
	
	homeDir, _ := os.UserHomeDir()
	fullPath := filepath.Join(homeDir, file)
	jsonFile, err := os.Open(fullPath)
	
	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteVal, err := io.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(byteVal, &cfg)
	if err != nil {
		fmt.Println(err)
	}

	return cfg
}

func (c *Config) SetUser() {

}
