package cli

import (
	"fmt"
	"github.com/mxlang/dotx/internal/config"
	"github.com/mxlang/dotx/internal/git"
	"github.com/mxlang/dotx/internal/logger"
	"github.com/spf13/cobra"
)

func newCmdStatus(cfg *config.Config) *cobra.Command {
	var prompt bool

	statusCmd := &cobra.Command{
		Use:   "status",
		Short: "Show status of dotfiles repository",
		Long:  "Show if your dotfiles repository is up to date with the remote",
		Example: `  dotx status
  dotx status --prompt"`,

		Args: cobra.NoArgs,

		Run: func(cmd *cobra.Command, args []string) {
			isBehind, err := git.IsBehindRemote(cfg.RepoPath)
			if err != nil {
				logger.Error("failed to get status of dotfiles repository", "error", err)
			}

			if prompt {
				fmt.Println(isBehind)
				return
			}

			if isBehind {
				logger.Info("remote dotfiles have changes")
			}
		},
	}

	statusCmd.PersistentFlags().BoolVarP(&prompt, "prompt", "p", false, "return boolean used for prompts")

	return statusCmd
}
