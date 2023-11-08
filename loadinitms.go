package loadinitms

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

// loadProperties loads the properties from a file.
func loadProperties() {
	LoadProperties()
}

// runPrimaries executes the primary processes.
func runPrimaries() {
	RunPrimaries()
}

// Run starts LoadInitMS, loads properties, executes primary processes, and handles exit signals.
func Run() {
	loadProperties()
	runPrimaries()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGABRT, syscall.SIGTERM)
	s := <-signalChan
	log.Printf("Exit signal activated: %v", s)
	os.Exit(0)
}
