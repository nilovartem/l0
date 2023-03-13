package main

import (
	"fmt"
	"net/url"
	"os"
	"path/filepath"

	"github.com/nats-io/stan.go"
	"github.com/nilovartem/l0/cmd/config"
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
	getConfig()
	fmt.Println("Hello, i am publisher!")
	filename, _ := filepath.Abs("../publisher/model.json")
	contents, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("File reading error", err)
		panic(err)
	}
	sc, err := stan.Connect(cfg.STAN.ClusterID, cfg.STAN.ClientID, stan.NatsURL(getURL()))
	sc.Publish(cfg.STAN.Channel, []byte(contents))
	if err != nil {
		panic(err)
	}
}
