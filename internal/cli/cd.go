package cli

import (
	"fmt"
	"github.com/mxlang/dotx/internal/config"
	"github.com/spf13/cobra"
)

func newCmdInitShell() *cobra.Command {
	return &cobra.Command{
		Use:    "init",
		Hidden: true,

		Args: cobra.NoArgs,

		Run: func(cmd *cobra.Command, args []string) {
			fmt.Print(`
dotx() {
  if [ "$1" = "cd" ]; then
    shift
    cd "$(command dotx cd "$@")"
  else
    command dotx "$@"
  fi
}
`)
		},
	}
}

func newCmdCd(cfg *config.Config) *cobra.Command {
	return &cobra.Command{
		Use:     "cd",
		Short:   "",
		Long:    "",
		Example: "  dotx cd",

		Args: cobra.NoArgs,

		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(cfg.RepoPath)
		},
	}
}
