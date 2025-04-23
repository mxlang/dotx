package cli

import (
	"github.com/mxlang/dotx/internal/dotx"
	"github.com/spf13/cobra"
)

func newCmdSync(dotx dotx.App) *cobra.Command {
	syncCmd := &cobra.Command{
		Use:   "sync",
		Short: "Manage every git related operations for your dotfiles repository",
		Args:  cobra.NoArgs,
	}

	syncCmd.AddCommand(
		newCmdInit(dotx),
		newCmdPull(dotx),
		newCmdPush(dotx),
	)

	return syncCmd
}
