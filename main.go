package main

import (
	"wb0-app/logger"
)

func main() {
	logg := logger.Setup()
	logg.Debug("Hello world!")
}
