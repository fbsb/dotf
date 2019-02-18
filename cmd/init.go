package cmd

import (
	"fmt"

	"github.com/fbsb/dotf/repo"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize directory for storing local configuration files.",
	RunE: func(cmd *cobra.Command, args []string) error {
		_, err := repo.Init(rootPath)
		if err != nil {
			return err
		}

		fmt.Println("Successfully initialized dotfiles repository in", rootPath)

		return nil
	},
}

func init() {
	dotfCmd.AddCommand(initCmd)
}
