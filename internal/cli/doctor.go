package cli

import (
	"github.com/mxlang/dotx/internal/core"
	"github.com/mxlang/dotx/internal/logger"
	"github.com/spf13/cobra"
)

func newCmdDoctor(app core.App) *cobra.Command {
	var force bool

	doctorCmd := &cobra.Command{
		Use:     "doctor",
		Short:   "",
		Long:    "",
		Example: `  `,

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
