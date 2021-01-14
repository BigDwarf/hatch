package main

import (
	"github.com/sirupsen/logrus"
	"hatch/main/comparators"
	"math/rand"
	"os"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	logrus.SetFormatter(&logrus.TextFormatter{TimestampFormat: "2021-01-13 19:04:05.999", FullTimestamp: true})

	command := comparators.NewCompareCommand()

	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
