package cli

import (
	"github.com/mxlang/dotx/internal/dotx"
	"github.com/spf13/cobra"
)

func newCmdSync(dotx dotx.App) *cobra.Command {
	syncCmd := &cobra.Command{
		Use:   "sync",
		Short: "Manage Git operations for your dotfiles repository",
		Long:  "Perform Git operations such as initializing, pulling, and pushing changes to synchronize your dotfiles across systems",
		Example: `  dotx sync init https://github.com/username/dotfiles.git
  dotx sync pull
  dotx sync push -m "Update nvim configuration"`,

		Args: cobra.NoArgs,
	}

	syncCmd.AddCommand(
		newCmdInit(dotx),
		newCmdPull(dotx),
		newCmdPush(dotx),
	)

	return syncCmd
}
