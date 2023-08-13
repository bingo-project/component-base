package options

import (
	"github.com/spf13/cobra"

	"github.com/bingo-project/component-base/cli/templates"
)

// NewCmdOptions returns new initialized instance of 'options' sub command.
func NewCmdOptions() *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "options",
		DisableFlagsInUseLine: true,
		Short:                 "Print the list of flags inherited by all commands",
		TraverseChildren:      true,
		Run: func(cmd *cobra.Command, args []string) {
			_ = cmd.Usage()
		},
	}

	templates.UseOptionsTemplates(cmd)

	return cmd
}
