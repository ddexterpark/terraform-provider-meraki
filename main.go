package main

import (
	"context"
	"flag"
	"github.com/ddexterpark/terraform-provider-meraki/internal/provider"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"log"
)

// Run "go generate" to format example terraform files and generate the docs for the registry/website

// If you do not have terraform installed, you can remove the formatting command, but its suggested to
// ensure the documentation is formatted properly.
//go:generate terraform fmt -recursive ./examples/

// Run the docs generation tool, check its repository for more information on how it works and how docs
// can be customized.
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs

var (
	// these will be set by the goreleaser configuration
	// to appropriate values for the compiled binary
	version string = "dev"

	// goreleaser can also pass the specific commit if you want
	// commit  string = ""
)

// Terraform has detailed logs which can be enabled by setting the TF_LOG environment variable to any value. This will cause detailed logs to appear on stderr.
//You can set TF_LOG to one of the log levels TRACE, DEBUG, INFO, WARN or ERROR to change the verbosity of the logs.
//Setting TF_LOG to JSON outputs logs at the TRACE level or higher, and uses a parseable JSON encoding as the formatting.

// export TF_LOG=DEBUG

// Redirect stderr to file instead of viewing in console
// export TF_LOG_PATH=${PWD}/logs.txt

func main() {
	var debug bool

	flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	opts := tfsdk.ServeOpts{
		Debug: debug,

		// TODO: Update this string with the published name of your provider.
		Name: "meraki",
	}

	err := tfsdk.Serve(context.Background(), provider.New(version), opts)

	if err != nil {
		log.Fatal(err.Error())
	}

}
