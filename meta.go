package main

import (
	"flag"
	"github.com/duaraghav8/okta-admin/common"
	"log"
	"os"
)

// createMeta returns the Metadata object to be passed to
// all actions. None of the global options are treated as
// required. Checking for emptiness of an option and
// further validations are therefore the user's responsibility.
func createMeta(logger *log.Logger) (*common.CommandMetadata, error) {
	var (
		meta       common.CommandMetadata
		globalOpts common.CommandConfig
	)
	flags := flag.NewFlagSet("options", flag.ContinueOnError)

	flags.StringVar(&globalOpts.OrgUrl, "org-url", os.Getenv("OKTA_ORG_URL"), "")
	flags.StringVar(&globalOpts.ApiToken, "api-token", os.Getenv("OKTA_API_TOKEN"), "")

	meta = common.CommandMetadata{
		Logger:        logger,
		FlagSet:       flags,
		GlobalOptions: &globalOpts,
		GlobalOptionsHelpText: `
Global Options:
  -org-url   Okta organization URL
             This can also be specified via the OKTA_ORG_URL environment variable.
  -api-token Token to authenticate with Okta API
             This can also be specified via the OKTA_API_TOKEN environment variable.
`,
	}

	return &meta, nil
}
