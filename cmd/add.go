package cmd

import (
	"log"

	"github.com/mlang97/dotx/config"
	"github.com/spf13/cobra"
)

func newCmdAdd(cfg config.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "add",
		Short: "Add file to your dotfiles repository",
		Args:  cobra.ExactArgs(1),

		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			if err := cfg.CheckDotfilesDir(); err != nil {
				return err
			}
			cfg.EnsureRepoConfig()
			return nil
		},

		Run: func(cmd *cobra.Command, args []string) {
			if err := cfg.AddDotfile(args[0]); err != nil {
				log.Fatal(err)
			}
		},
	}
}
