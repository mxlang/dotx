package cmd

import (
	"os"

	"github.com/mlang97/dotx/app"
	"github.com/spf13/cobra"
)

func Execute(dotx app.App) {
	err := newCmdRoot(dotx).Execute()
	if err != nil {
		os.Exit(1)
	}
}

func newCmdRoot(dotx app.App) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "dotx",
		Short: "The next generation dotfile manager",
		Args:  cobra.NoArgs,

		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			dotx.EnsureRepo()
		},
	}

	rootCmd.AddCommand(
		newCmdAdd(dotx),
		newCmdDeploy(dotx),
	)

	return rootCmd
}
