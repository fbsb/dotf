package cmd

import (
	"fmt"

	"github.com/fbsb/dotf/repo/git"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [file1...]",
	Short: "Add a file to the repository and track changes made to it.",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		r, err := git.Open(rootPath)
		if err != nil {
			return err
		}

		for _, arg := range args {
			err := r.Add(arg)
			if err != nil {
				return err
			}
			fmt.Println("added", arg)
		}

		return nil
	},
}

func init() {
	dotfCmd.AddCommand(addCmd)
}
