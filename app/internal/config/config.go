package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"time"
)

type Config struct {
	Server   `yaml:"server"`
	Postgres `yaml:"postgres"`
	Mongo    `yaml:"mongo"`
}

type Server struct {
	Address string        `yaml:"address"`
	Timeout time.Duration `yaml:"timeout"`
}

type Postgres struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	DBName   string `yaml:"dbname"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

type Mongo struct {
}

func MustLoad() Config {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("config path is not set")
	}
	var config Config
	err := cleanenv.ReadConfig(configPath, &config)
	if err != nil {
		log.Fatalf("config not read: %v", err)
	}
	return config
}
