package cmd

import (
	"fmt"

	"github.com/fbsb/dotf/repo"
	"github.com/spf13/cobra"
)

var untrackCmd = &cobra.Command{
	Use:   "untrack [file1...]",
	Short: "Untrack one or more files or directories.",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		r, err := repo.Open(rootPath)
		if err != nil {
			return err
		}

		for _, arg := range args {
			err := r.Untrack(arg)
			if err != nil {
				return err
			}
			fmt.Println("Untracked", arg)
		}

		return nil
	},
}

func init() {
	dotfCmd.AddCommand(untrackCmd)
}
