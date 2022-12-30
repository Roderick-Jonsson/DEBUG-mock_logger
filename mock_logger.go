package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {
	println("=== STARTING APP MOCK LOGGER ===")
	log.Trace("Something very low level.")
	log.Debug("Useful debugging information.")
	log.Info("Something noteworthy happened!")
	log.Warn("You should probably take a look at this.")
	log.Error("Something failed but I'm not quitting.")
	println("=== ENDING APP MOCK LOGGER ===")
}
