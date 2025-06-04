package main

import (
	"go.uber.org/zap"
)

var log = zap.Must(zap.NewDevelopment())

func main() {
	defer func() { _ = log.Sync() }()
	log.Info("Hello World")
}
