package cli

import (
	"fmt"
	"strings"

	"github.com/mxlang/dotx/internal/config"
	"github.com/mxlang/dotx/internal/logger"
	"github.com/spf13/cobra"
)

func newCmdInitShell() *cobra.Command {
	return &cobra.Command{
		Use:    "init",
		Hidden: true,

		ValidArgs: []string{"bash", "zsh", "fish"},
		Args:      cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),

		Run: func(cmd *cobra.Command, args []string) {
			shell := args[0]

			w := cmd.OutOrStdout()
			posixFunction := "\n" + `dotx() { if [ "$1" = "cd" ]; then cd "$(command dotx cd)"; else command dotx "$@"; fi; }` + "\n"

			switch strings.ToLower(shell) {
			case "bash":
				w.Write([]byte(posixFunction))
				cmd.Root().GenBashCompletion(w)
			case "zsh":
				w.Write([]byte(posixFunction))
				cmd.Root().GenZshCompletion(w)
			case "fish":
				w.Write([]byte("\n" + `function dotx; if test (count $argv) -ge 1; and test $argv[1] = "cd"; cd (command dotx cd); else; command dotx $argv; end; end` + "\n"))
				cmd.Root().GenFishCompletion(w, true)
			default:
				logger.Error("unsupported shell", "shell", shell)
			}
		},
	}
}

func newCmdCd(cfg *config.Config) *cobra.Command {
	return &cobra.Command{
		Use:     "cd",
		Short:   "Go to your local dotfiles directory",
		Example: "  dotx cd",

		Args: cobra.NoArgs,

		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(cfg.RepoPath)
		},
	}
}
