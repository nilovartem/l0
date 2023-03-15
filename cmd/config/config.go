package config

import (
	"os"
	"path/filepath"

	yaml "gopkg.in/yaml.v3"
)

type Config struct {
	STAN struct {
		Scheme      string `yaml:"scheme"`
		Host        string `yaml:"host"`
		Port        string `yaml:"port"`
		ClusterID   string `yaml:"clusterID"`
		PublisherID string `yaml:"publisherID"`
		ConsumerID  string `yaml:"consumerID"`
		Channel     string `yaml:"channel"`
	} `yaml:"STAN"`
	Cache struct {
		ModelKey string `yaml:"modelKey"`
	} `yaml:"Cache"`
}

func (config *Config) GetConfig() {
	filename, _ := filepath.Abs("../config/env.yaml")
	contents, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(contents, &config)
	if err != nil {
		panic(err)
	}
	//TODO:return error
}
