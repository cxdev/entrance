package main

import (
	entrance "entrance/backend"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type AppConfig struct {
	Port    int    `yaml:"port"`
	RootDir string `yaml:"root_dir"`
}

type CommandConfig entrance.Command

func loadYaml(path string, object interface{}) error {
	configData, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	log.Printf("Loading config:\n%s", configData)

	return yaml.Unmarshal(configData, object)
}

func loadAppConfig(configPath string) (*AppConfig, error) {
	var config AppConfig
	err := loadYaml(configPath, &config)
	return &config, err
}

func loadCommandConfig(configPath string) (*CommandConfig, error) {
	var config CommandConfig
	err := loadYaml(configPath, &config)
	return &config, err
}
