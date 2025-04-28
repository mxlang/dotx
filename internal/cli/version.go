package cli

import (
	"fmt"
	"github.com/mxlang/dotx/internal/dotx"
	"github.com/spf13/cobra"
)

func newCmdVersion(dotx dotx.App) *cobra.Command {
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Print dotx version",
		Long:  "Display the current version of dotx, which helps track which features and bug fixes are available in your installation",

		Args: cobra.NoArgs,

		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("dotx version 0.1.0")
		},
	}

	return versionCmd
}
