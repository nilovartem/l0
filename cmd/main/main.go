package main

import (
	"github.com/nilovartem/l0/cmd/consumer"
	"github.com/nilovartem/l0/cmd/memory"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Infoln("Starting main")
	memory.New()
	consumer.Run()
}
