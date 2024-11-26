package config

import (
	"errors"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
	"time"
)

var Config struct {
	Server struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"server"`

	Database struct {
		Driver                string        `yaml:"driver"`
		Host                  string        `yaml:"host"`
		Port                  string        `yaml:"port"`
		Username              string        `yaml:"username"`
		Password              string        `yaml:"password"`
		DbName                string        `yaml:"name"`
		Sslmode               string        `yaml:"sslmode"`
		MaxConnections        int           `yaml:"max_connections"`
		MaxIdleConnections    int           `yaml:"max_idle_connections"`
		ConnectionMaxLifetime time.Duration `yaml:"connection_max_lifetime"`
	} `yaml:"database"`

	Cache struct {
		Type  string `yaml:"type"`
		Redis struct {
			Host     string        `yaml:"host"`
			Port     int           `yaml:"port"`
			Password string        `yaml:"password"`
			DB       int           `yaml:"db"`
			TTL      time.Duration `yaml:"ttl"`
		} `yaml:"redis"`
		Memory struct {
			DefaultExpiration time.Duration `yaml:"default_expiration"`
			CleanupInterval   time.Duration `yaml:"cleanup_interval"`
		} `yaml:"memory"`
	} `yaml:"cache"`

	JWT struct {
		Secret          string        `yaml:"secret"`
		AccessTokenTTL  time.Duration `yaml:"access_token_ttl"`
		RefreshTokenTTL time.Duration `yaml:"refresh_token_ttl"`
	} `yaml:"jwt"`
}

func LoadConfig() error {
	basePath, err := os.Getwd()
	if err != nil {
		return errors.New("failed to get current directory: " + err.Error())
	}

	configPath := filepath.Join(basePath, "auth", "pkg", "config", "config.yaml")

	file, err := os.Open(configPath)
	if err != nil {
		return errors.New("failed to open config file: " + err.Error())
	}

	defer file.Close()

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&Config); err != nil {
		return errors.New("failed to decode config file: " + err.Error())
	}

	return nil
}
