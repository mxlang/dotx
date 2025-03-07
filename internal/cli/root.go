package cli

import (
	"os"

	"github.com/mlang97/dotx/internal/dotx"
	"github.com/spf13/cobra"
)

func Execute(dotx dotx.App) {
	err := newCmdRoot(dotx).Execute()
	if err != nil {
		os.Exit(1)
	}
}

func newCmdRoot(dotx dotx.App) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "dotx",
		Short: "The next generation dotfile manager",
		Args:  cobra.NoArgs,
	}

	rootCmd.AddCommand(
		newCmdAdd(dotx),
		newCmdDeploy(dotx),
		newCmdSync(dotx),
	)

	return rootCmd
}
