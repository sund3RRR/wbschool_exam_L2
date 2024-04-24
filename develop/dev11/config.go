package main

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Port     int    `yaml:"port"`
	DataPath string `yaml:"data_path"`
}

func NewConfig() *Config {
	data, err := os.ReadFile("config.yml")
	if err != nil {
		log.Fatalf("Error reading config.yml: %v", err)
	}

	cfg := Config{}

	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		log.Fatalf("Error parsing YAML: %v", err)
	}

	return &cfg
}
