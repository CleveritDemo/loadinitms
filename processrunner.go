package loadinitms

import (
	"log"
)

// PrimaryProcess is the primary process to be run by the load-init-ms library
type PrimaryProcess interface {
	Start()
}

var primaries = make([]PrimaryProcess, 0)

// AddPrimary adds a primary process to the list of primaries
func AddPrimary(primary PrimaryProcess) {
	primaries = append(primaries, primary)
}

// runPrimaries runs all the primaries concurrently
func runPrimaries() {
	done := make(chan struct{})

	for _, primary := range primaries {
		go func(p PrimaryProcess) {
			log.Printf("Running primary process: %T\n", p)
			p.Start()
			done <- struct{}{}
		}(primary)
	}

	// Wait for all primaries to complete
	for range primaries {
		<-done
	}
}
