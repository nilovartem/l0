package main

import (
	"fmt"
	"net/url"
	"os"
	"path/filepath"

	"github.com/nats-io/stan.go"
	"github.com/nilovartem/l0/cmd/config"
	"github.com/sirupsen/logrus"
)

var (
	cfg config.Config
)

func getURL() string {
	url := url.URL{
		Scheme: cfg.STAN.Scheme,
		Host:   cfg.STAN.Host + cfg.STAN.Port,
	}
	return url.String()
}
func getConfig() { cfg.GetConfig() }

func main() {
	fmt.Println("Starting publisher")
	getConfig()
	filename := os.Args[1]
	filename, _ = filepath.Abs("../publisher/data/" + filename)
	contents, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	sc, err := stan.Connect(cfg.STAN.ClusterID, cfg.STAN.PublisherID, stan.NatsURL(getURL()))
	if err != nil {
		logrus.Errorln("[FAIL] Cannot connect to STAN, aborting")
		logrus.Errorln(err)
		return
	}
	err = sc.Publish(cfg.STAN.Channel, []byte(contents))
	if err != nil {
		logrus.Errorln("[FAIL] Cannot publish to channel, aborting")
		logrus.Errorln(err)
	}
}
