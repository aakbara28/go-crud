package environment

import (
	"encoding/json"
	"fmt"
	"os"
)

type DatabaseConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Name     string `json:"name"`
	SSLMode  string `json:"sslmode"`
}

type Config struct {
	Database DatabaseConfig `json:"database"`
	Server   ServerConfig   `json:"server"`
}

type ServerConfig struct {
	Port string `json:"port"`
}

var AppConfig Config

func LoadConfig(env string) error {
	filename := fmt.Sprintf("%s.json", env)
	
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&AppConfig)
	if err != nil {
		return err
	}
	
	return nil
}