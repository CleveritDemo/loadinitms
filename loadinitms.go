package loadinitms

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

// LoadInitMS is the structure that encapsulates the functionality of LoadInitMS.
type LoadInitMS struct{}

// NewLoadInitMS creates a new instance of LoadInitMS.
func NewLoadInitMS() *LoadInitMS {
	return &LoadInitMS{}
}

// loadProperties loads the properties from a file.
func (lm *LoadInitMS) loadProperties() {
	LoadProperties()
}

// runPrimaries executes the primary processes.
func (lm *LoadInitMS) runPrimaries() {
	RunPrimaries()
}

// Run starts LoadInitMS, loads properties, executes primary processes, and handles exit signals.
func (lm *LoadInitMS) Run() {
	lm.loadProperties()
	lm.runPrimaries()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGABRT, syscall.SIGTERM)
	s := <-signalChan
	log.Printf("Exit signal activated: %v", s)
	os.Exit(0)
}
