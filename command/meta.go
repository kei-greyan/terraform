package command

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/mitchellh/cli"
	"github.com/mitchellh/colorstring"
)

// Meta contains the meta-options and functionality that nearly every
// Terraform command inherits.
type Meta struct {
	// Color is set to true if colors should be enabled.
	Color bool

	// GlobalPluginDirs is a list of directories to search for plugins.
	GlobalPluginDirs []string

	// PluginPath is a list of paths to search for plugins specified by the user
	// via the -plugin-dir flag.
	PluginPath []string

	// Ui is the user interface for input/output.
	Ui cli.Ui

	// input is whether input is enabled.
	input bool

	// autoApprove is whether the user has approved all prompts.
	autoApprove bool

	// workspace is the workspace to use.
	workspace string

	// color is the colorize object to use for colorized output.
	color *colorstring.Colorize
}

// defaultParallelism is the default number of parallel operations.
const defaultParallelism = 10

// flagSet adds the meta flags to the given flag set.
func (m *Meta) flagSet(n string) *flag.FlagSet {
	f := flag.NewFlagSet(n, flag.ContinueOnError)
	f.BoolVar(&m.input, "input", true, "Ask for input for variables if not directly set.")
	f.BoolVar(&m.Color, "no-color", false, "If specified, output won't contain any color.")
	return f
}

// colorize returns the colorize object to use for output.
func (m *Meta) colorize() *colorstring.Colorize {
	if m.color != nil {
		return m.color
	}

	m.color = &colorstring.Colorize{
		Colors:  colorstring.DefaultColors,
		Disable: !m.Color,
		Reset:   true,
	}
	return m.color
}

// outputColumns returns the number of columns for output.
func (m *Meta) outputColumns() int {
	if cols := os.Getenv("COLUMNS"); cols != "" {
		var n int
		if _, err := fmt.Sscanf(cols, "%d", &n); err == nil && n > 0 {
			return n
		}
	}
	return 78
}

// workspace returns the name of the currently configured workspace.
func (m *Meta) Workspace() string {
	if m.workspace != "" {
		return m.workspace
	}
	if ws := os.Getenv("TF_WORKSPACE"); ws != "" {
		return ws
	}
	return "default"
}

// confirm asks the user for confirmation of the given message. It returns
// true if the user confirms, false otherwise.
func (m *Meta) confirm(query string) (bool, error) {
	if !m.input {
		return false, fmt.Errorf("input is disabled, cannot prompt for confirmation")
	}

	v, err := m.Ui.Ask(query)
	if err != nil {
		return false, fmt.Errorf("error asking for confirmation: %w", err)
	}

	switch v {
	case "yes":
		return true, nil
	default:
		return false, nil
	}
}

// stdinPiped returns true if stdin is piped (non-interactive).
func (m *Meta) stdinPiped() bool {
	stat, err := os.Stdin.Stat()
	if err != nil {
		return false
	}
	return (stat.Mode() & os.ModeCharDevice) == 0
}

// lineWriter is an io.Writer that writes to the given Ui line by line.
type lineWriter struct {
	ui     cli.Ui
	buf    *bufio.Writer
	writer io.Writer
}
