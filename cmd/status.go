package cmd

import (
	"fmt"

	"github.com/fbsb/dotf/repo"
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show status of tracked configuration files",
	RunE: func(cmd *cobra.Command, args []string) error {
		r, err := repo.Open(rootPath)
		if err != nil {
			return err
		}

		s, err := r.Status()
		if err != nil {
			return err
		}

		if len(s) > 0 {
			fmt.Println("Changed files:")
			fmt.Println(s)
		} else {
			fmt.Println("No changes.")
		}

		return nil
	},
}

func init() {
	dotfCmd.AddCommand(statusCmd)
}
