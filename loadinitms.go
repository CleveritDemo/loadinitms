package loadinitms

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

// ExecLoadInitMS executes LoadInitMS, loads properties, and executes primary processes.
func ExecLoadInitMS() {
	loadProperties()
	runPrimaries()
}

// Run starts LoadInitMS, loads properties, executes primary processes, and handles exit signals.
func Run() {
	ExecLoadInitMS()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGABRT, syscall.SIGTERM)
	s := <-signalChan
	log.Printf("Exit signal activated: %v", s)
	os.Exit(0)
}
