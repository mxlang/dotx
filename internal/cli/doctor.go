package cli

import (
	"github.com/mxlang/dotx/internal/core"
	"github.com/mxlang/dotx/internal/logger"
	"github.com/spf13/cobra"
)

func newCmdDoctor(app core.App) *cobra.Command {
	var force bool

	doctorCmd := &cobra.Command{
		Use:   "doctor",
		Short: "Scan for unlinked dotfiles and offer to deploy them",
		Long:  "Inspect your dotx repository and current system for dotfiles that exist in the repo but are not yet deployed (symlinked). Choose which ones to link. Use --force to never prompt when overwriting.",
		Example: `  dotx doctor
  dotx doctor --force`,

		Args: cobra.NoArgs,

		Run: func(cmd *cobra.Command, args []string) {
			if err := app.Doctor(force); err != nil {
				logger.Error("failed to run doctor", "error", err)
			}
		},
	}

	doctorCmd.PersistentFlags().BoolVarP(&force, "force", "f", false, "never prompt for overwriting")

	return doctorCmd
}
