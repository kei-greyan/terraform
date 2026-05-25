package main

import (
	"os"

	"github.com/mitchellh/cli"

	"github.com/opentofu/opentofu/internal/command"
	"github.com/opentofu/opentofu/internal/command/meta"
)

// Commands is the mapping of all the available Terraform commands.
var Commands map[string]cli.CommandFactory

// PlumbingCommands is the list of commands that are considered
// "plumbing" and not user-facing. These commands are not shown
// in the main help output.
var PlumbingCommands map[string]struct{}

// initCommands initializes the Commands and PlumbingCommands variables.
// It is called from main() to set up the CLI commands.
func initCommands(
	ctx context.Context,
	originalWorkingDir string,
	streams *terminal.Streams,
	config *cliconfig.Config,
	services *disco.Disco,
	providerSrc getproviders.Source,
	providerDevOverrides map[addrs.Provider]getproviders.PackageLocalDir,
	unmanagedProviders map[addrs.Provider]*plugin.ReattachConfig,
) {
	var inAutomation bool
	if v := os.Getenv("TF_IN_AUTOMATION"); v != "" {
		inAutomation = true
	}

	workspaceNameEnvVar := os.Getenv("TF_WORKSPACE")

	PlumbingCommands = map[string]struct{}{
		"state": {},
		"debug": {},
		"force-unlock": {},
		"push": {},
		"0.12upgrade": {},
		"0.13upgrade": {},
	}

	meta := meta.Meta{
		OriginalWorkingDir: originalWorkingDir,
		Streams:            streams,
		Color:              true,
		GlobalPluginsDir:   globalPluginDirs(),
		AutomationReason:   automationReason(inAutomation),
		WorkspaceNameEnvVar: workspaceNameEnvVar,
		Services:           services,
		ProviderSource:     providerSrc,
		ProviderDevOverrides: providerDevOverrides,
		UnmanagedProviders: unmanagedProviders,
	}

	Commands = map[string]cli.CommandFactory{
		"apply": func() (cli.Command, error) {
			return &command.ApplyCommand{
				Meta: meta,
			}, nil
		},
		"console": func() (cli.Command, error) {
			return &command.ConsoleCommand{
				Meta: meta,
			}, nil
		},
		"destroy": func() (cli.Command, error) {
			return &command.ApplyCommand{
				Meta:    meta,
				Destroy: true,
			}, nil
		},
		"fmt": func() (cli.Command, error) {
			return &command.FmtCommand{
				Meta: meta,
			}, nil
		},
		"get": func() (cli.Command, error) {
			return &command.GetCommand{
				Meta: meta,
			}, nil
		},
		"graph": func() (cli.Command, error) {
			return &command.GraphCommand{
				Meta: meta,
			}, nil
		},
		"import": func() (cli.Command, error) {
			return &command.ImportCommand{
				Meta: meta,
			}, nil
		},
		"init": func() (cli.Command, error) {
			return &command.InitCommand{
				Meta: meta,
			}, nil
		},
		"login": func() (cli.Command, error) {
			return &command.LoginCommand{
				Meta: meta,
			}, nil
		},
		"logout": func() (cli.Command, error) {
			return &command.LogoutCommand{
				Meta: meta,
			}, nil
		},
		"output": func() (cli.Command, error) {
			return &command.OutputCommand{
				Meta: meta,
			}, nil
		},
		"plan": func() (cli.Command, error) {
			return &command.PlanCommand{
				Meta: meta,
			}, nil
		},
		"providers": func() (cli.Command, error) {
			return &command.ProvidersCommand{
				Meta: meta,
			}, nil
		},
		"refresh": func() (cli.Command, error) {
			return &command.RefreshCommand{
				Meta: meta,
			}, nil
		},
		"show": func() (cli.Command, error) {
			return &command.ShowCommand{
				Meta: meta,
			}, nil
		},
		"taint": func() (cli.Command, error) {
			return &command.TaintCommand{
				Meta: meta,
			}, nil
		},
		"untaint": func() (cli.Command, error) {
			return &command.UntaintCommand{
				Meta: meta,
			}, nil
		},
		"validate": func() (cli.Command, error) {
			return &command.ValidateCommand{
				Meta: meta,
			}, nil
		},
		"version": func() (cli.Command, error) {
			return &command.VersionCommand{
				Meta:              meta,
				Version:           Version,
				VersionPrerelease: VersionPrerelease,
				VersionMetadata:   VersionMetadata,
				CheckFunc:         commandVersionCheck,
			}, nil
		},
		"workspace": func() (cli.Command, error) {
			return &command.WorkspaceCommand{
				Meta: meta,
			}, nil
		},
	}
}

// automationReason returns a human-readable string describing why
// Terraform is running in automation mode, or an empty string if it is not.
func automationReason(inAutomation bool) string {
	if !inAutomation {
		return ""
	}
	return "TF_IN_AUTOMATION is set"
}
