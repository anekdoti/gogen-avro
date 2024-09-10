package compiler

import (
	"fmt"
)

var (
	// Enable this to get debug logs for the compilation process
	LoggingEnabled = true
)

func log(f string, v ...interface{}) {
	if LoggingEnabled {
		fmt.Printf(f+"\n", v...)
	}
}
