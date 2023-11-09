package loadinitms

import (
	"log"
)

// PrimaryProcess is the primary process to be run by the load-init-ms library
type PrimaryProcess interface {
	Start()
}

// primaryProcessFunc is a function that returns a PrimaryProcess
type primaryProcessFunc func() PrimaryProcess

var primaries = make([]func() PrimaryProcess, 0)

// AddPrimary adds a primary process to the list of primaries
func AddPrimary(primary primaryProcessFunc) {
	primaries = append(primaries, primary)
}

// runPrimaries runs all the primaries concurrently
func runPrimaries() {
	for _, primary := range primaries {
		log.Printf("Running primary process: %T\n", primary())
		go primary().Start()
	}
}
