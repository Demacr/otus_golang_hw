package config

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/spf13/pflag"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Host       string `yaml:"host"`
	Port       int    `yaml:"port"`
	Database   string `yaml:"database"`
	PostgreSQL struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Login    string `yaml:"login"`
		Password string `yaml:"password"`
		Database string `yaml:"database"`
	} `yaml:"postgresql"`
	Log struct {
		File  string `yaml:"file"`
		Level string `yaml:"level"`
	} `yaml:"log"`
}

func Configure() *Config {
	var configPath string
	var config Config
	pflag.StringVar(&configPath, "config", "", "Path to config file")
	pflag.Parse()

	if configPath == "" {
		log.Fatal("missing config file")
	}

	configFile, err := os.Open(configPath)
	if err != nil {
		log.Fatal(err)
	}

	configBytes, err := ioutil.ReadAll(configFile)
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(configBytes, &config)
	if err != nil {
		log.Fatal(err)
	}

	return &config
}
