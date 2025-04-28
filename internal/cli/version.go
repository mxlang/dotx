package cli

import (
	"fmt"
	"github.com/spf13/cobra"
)

func newCmdVersion(version string) *cobra.Command {
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Print dotx version",
		Long:  "Display the current version of dotx, which helps track which features and bug fixes are available in your installation",

		Args: cobra.NoArgs,

		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("dotx version %s\n", version)
		},
	}

	return versionCmd
}
