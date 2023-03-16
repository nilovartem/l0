package config

import (
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
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
	DB struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		DBname   string `yaml:"dbname"`
	} `yaml:"DB"`
}

/*Парсим в структур конфигурационный файл env.yaml*/
func (config *Config) GetConfig() {
	filename, _ := filepath.Abs("../config/env.yaml")
	contents, err := os.ReadFile(filename)
	if err != nil {
		logrus.Errorln("[FAIL] Can't read config file, aborting")
		panic(err)
	}
	err = yaml.Unmarshal(contents, &config)
	if err != nil {
		logrus.Errorln("[FAIL] Can't unmarshal config file, aborting")
		panic(err)
	}
}
