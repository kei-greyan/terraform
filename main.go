// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

// Terraform is a tool for building, changing, and versioning infrastructure
// safely and efficiently. Configuration files describe to Terraform the
// components needed to run a single application or your entire datacenter.
// Terraform generates an execution plan describing what it will do to reach
// the desired state, and then executes it to build the described
// infrastructure. As the configuration changes, Terraform is able to determine
// what changed and create incremental execution plans which can be applied.
//
// For more information, see the README and documentation in the docs/ directory.
package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/opentofu/opentofu/internal/command/meta"
	"github.com/opentofu/opentofu/version"
)

func main() {
	// Call realMain so that deferred functions work properly, since os.Exit
	// does not execute deferred functions.
	os.Exit(realMain())
}

func realMain() int {
	defer panicHandler()

	// Set the Go garbage collector to use a target of 20% overhead, which
	// is more aggressive than the default of 100%. This helps keep memory
	// usage lower for long-running operations.
	if os.Getenv("GOGC") == "" {
		runtime.GC()
	}

	// NOTE: We're intentionally not using the "flag" package here because
	// we want to handle flags ourselves for better UX.
	args := os.Args[1:]

	// Handle the case where no arguments are provided.
	if len(args) == 0 {
		printUsage()
		return 1
	}

	// Initialize the meta object which holds shared state across commands.
	meta := meta.Meta{
		Color: true,
	}
	_ = meta

	// Dispatch to the appropriate subcommand.
	cmd := args[0]
	switch cmd {
	case "version", "-version", "--version":
		fmt.Printf("Terraform v%s\n", version.Version)
		if version.Prerelease != "" {
			fmt.Printf("Prerelease: %s\n", version.Prerelease)
		}
		return 0
	case "-help", "--help", "help":
		printUsage()
		return 0
	default:
		fmt.Fprintf(os.Stderr, "Error: unknown command %q\n", cmd)
		fmt.Fprintf(os.Stderr, "Run 'terraform -help' for usage.\n")
		return 1
	}
}

// printUsage prints the top-level usage information for the terraform CLI.
func printUsage() {
	fmt.Printf(`Usage: terraform [global options] <subcommand> [args]

The available commands for execution are listed below.
The primary workflow commands are given first, followed by
less common or more advanced commands.

Main commands:
  init          Prepare your working directory for other commands
  validate      Check whether the configuration is valid
  plan          Show changes required by the current configuration
  apply         Create or update infrastructure
  destroy       Destroy previously-created infrastructure

All other commands:
  console       Try Terraform expressions at an interactive command prompt
  fmt           Reformat your configuration in the standard style
  force-unlock  Release a stuck lock on the current workspace
  get           Install or upgrade remote Terraform modules
  graph         Generate a Graphviz graph of the steps in an operation
  import        Associate existing infrastructure with a Terraform resource
  login         Obtain and save credentials for a remote host
  logout        Remove locally-stored credentials for a remote host
  metadata      Metadata related commands
  output        Show output values from your root module
  providers     Show the providers required for this configuration
  refresh       Update the state to match remote systems
  show          Show the current state or a saved plan
  state         Advanced state management
  taint         Mark a resource instance as not fully functional
  test          Execute integration tests for a module
  untaint       Remove the 'tainted' state from a resource instance
  workspace     Workspace management

Global options (use these before the subcommand, if any):
  -chdir=DIR    Switch to a different working directory before executing the
                given subcommand.
  -help         Show this help output, or the help for a specified subcommand.
  -version      An alias for the "version" subcommand.
`)
}
