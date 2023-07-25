package main

import (
	"context"

	// Attempts to increase the OS file descriptors - Fail silently
	_ "github.com/khulnasoft-labs/fdmax/autofdmax"
	"github.com/khulnasoft-labs/gologger"
	"github.com/khulnasoft-labs/reveal/runner"
)

func main() {
	// Parse the command line flags and read config files
	options := runner.ParseOptions()

	newRunner, err := runner.NewRunner(options)
	if err != nil {
		gologger.Fatal().Msgf("Could not create runner: %s\n", err)
	}

	err = newRunner.Run(context.Background())
	if err != nil {
		gologger.Fatal().Msgf("Could not run enumeration: %s\n", err)
	}
	// close the runner
	if newRunner != nil {
		newRunner.Close()
	}
}
