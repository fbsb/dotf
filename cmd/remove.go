package cmd

import (
	"fmt"

	"github.com/fbsb/dotf/repo"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove [file1...]",
	Short: "Remove one or more files or directories from dotf (but not the filesystem).",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		r, err := repo.Open(rootPath)
		if err != nil {
			return err
		}

		for _, arg := range args {
			err := r.Remove(arg)
			if err != nil {
				return err
			}
			fmt.Println("Removed", arg)
		}

		return nil
	},
}

func init() {
	dotfCmd.AddCommand(removeCmd)
}
