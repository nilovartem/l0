package main

import (
	"github.com/nilovartem/l0/cmd/consumer"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Infoln("Starting main")
	consumer.Run()
}
