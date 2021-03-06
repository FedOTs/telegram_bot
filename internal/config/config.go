package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Database struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`
	DbName   string `json:"dbname"`
}

type Telegram struct {
	Token   string `json:"token"`
	WebHook string `json:"webhook"`
}

type Config struct {
	Database `json:"database"`
	Telegram `json:"telegram"`
}

func LoadConfiguration(file string) (Config, error) {
	var config Config
	configFile, err := os.ReadFile(file)

	if err != nil {
		return config, err
	}

	err = json.Unmarshal(configFile, &config)

	if err != nil {
		return config, err
	}

	return config, nil
}

func (c *Config) GetConnString() string {
	return fmt.Sprintf("server=%s;user id=%s;password=%s;database=%s", c.Database.Host, c.Database.User, c.Database.Password, c.Database.DbName)
}

func (c *Config) GetTgTokenString() string {
	return c.Telegram.Token
}

func (c *Config) GetTgWebHookString() string {
	return c.Telegram.WebHook
}
