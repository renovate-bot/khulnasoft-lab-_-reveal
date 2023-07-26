package main

import (
	"context"
	"fmt"

	"github.com/khulnasoft-labs/gologger"
	"github.com/khulnasoft-labs/reveal"
	"github.com/khulnasoft-labs/reveal/sources"
)

func main() {
	opts := reveal.Options{
		Agents:   []string{"shodan"},
		Queries:  []string{"ssl:'hackerone.com'"},
		Limit:    50,
		MaxRetry: 2,
		Timeout:  20,
	}

	u, err := reveal.New(&opts)
	if err != nil {
		panic(err)
	}

	allagents := u.AllAgents()
	gologger.Info().Msgf("Available reveal agents/sources :")
	for _, v := range allagents {
		fmt.Println(v)
	}

	fmt.Println("\n\n- Reveal Results:")
	result := func(result sources.Result) {
		fmt.Println(result.IpPort())
	}

	// Execute executes and returns a channel with all results
	// ch , err := u.Execute(context.Background())

	// Execute with Callback calls u.Execute() internally and abstracts channel handling logic
	if err := u.ExecuteWithCallback(context.TODO(), result); err != nil {
		panic(err)
	}
}
