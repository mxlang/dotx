package cmd

import (
	"github.com/mlang97/dotx/app"
	"github.com/spf13/cobra"
)

func newCmdSync(dotx app.App) *cobra.Command {
	syncCmd := &cobra.Command{
		Use:   "sync",
		Short: "A brief description of your command",
		Args:  cobra.NoArgs,
	}

	syncCmd.AddCommand(newCmdInit(dotx))

	return syncCmd
}
