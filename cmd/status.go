package cmd

import (
	"fmt"

	"github.com/fbsb/dotf/repo/git"
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show status of tracked configuration files",
	RunE: func(cmd *cobra.Command, args []string) error {
		r, err := git.Open(rootPath)
		if err != nil {
			return err
		}

		s, err := r.Status()
		if err != nil {
			return err
		}

		if s.IsClean() {
			fmt.Println("No changes.")
			return nil
		}

		fmt.Println("Changed files:")
		fmt.Println(s)

		return nil
	},
}

func init() {
	dotfCmd.AddCommand(statusCmd)
}
