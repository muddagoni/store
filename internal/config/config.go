package config

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Server struct {
		Port int    `yaml:"port"`
		Host string `yaml:"host"`
	} `yaml:"server"`
}

func Load(configFile string) (*Config, error) {
	config := &Config{}

	if _, err := os.Stat(configFile); err != nil {
		log.Printf("could not find config file: %v", err)
	}
	data, err := os.ReadFile(configFile)
	if err != nil {
		log.Printf("error reading file: %v", err)
		return nil, fmt.Errorf("error reading file %v", err)
	}

	if err = yaml.Unmarshal(data, &config); err != nil {
		log.Printf("error unmarshalling config %v", err)
		return nil, fmt.Errorf("error unmarshalling config %v", err)
	}

	return config, nil
}
