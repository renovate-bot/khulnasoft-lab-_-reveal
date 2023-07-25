package runner

import (
	"github.com/khulnasoft-labs/gologger"
	updateutils "github.com/khulnasoft-labs/utils/update"
)

const banner = `
  __  ______  _________ _   _____  _____
 / / / / __ \/ ___/ __ \ | / / _ \/ ___/
/ /_/ / / / / /__/ /_/ / |/ /  __/ /    
\__,_/_/ /_/\___/\____/|___/\___/_/
`

// Version is the current version of reveal
const version = `v1.0.5`

// showBanner is used to show the banner to the user
func showBanner() {
	gologger.Print().Msgf("%s\n", banner)
	gologger.Print().Msgf("\t\tkhulnasoft.com\n\n")
}

// GetUpdateCallback returns a callback function that updates reveal
func GetUpdateCallback() func() {
	return func() {
		showBanner()
		updateutils.GetUpdateToolCallback("reveal", version)()
	}
}
