package config

import (
	"errors"
	"gopkg.in/yaml.v3"
	"os"
	"time"
)

var Config struct {
	Server struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	} `yaml:"server"`

	Database struct {
		Driver                string        `yaml:"driver"`
		Host                  string        `yaml:"host"`
		Port                  int           `yaml:"port"`
		Username              string        `yaml:"username"`
		Password              string        `yaml:"password"`
		Name                  string        `yaml:"name"`
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
		Secret     string        `yaml:"secret"`
		Expiration time.Duration `yaml:"expiration"`
	} `yaml:"jwt"`
}

func LoadConfig() error {
	file, err := os.Open("config.yaml")
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
