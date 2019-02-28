package cmd

import (
	"errors"

	"fmt"

	"github.com/fbsb/dotf/repo/git"
	"github.com/spf13/cobra"
)

var (
	message, name, email string
)

var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "Commit changes to tracked files.",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if len(message) == 0 {
			return errors.New("message parameter is requied")
		}

		if len(name) == 0 {
			return errors.New("name parameter is required")
		}

		if len(email) == 0 {
			return errors.New("email parameter is required")
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		r, err := git.Open(rootPath)
		if err != nil {
			return err
		}

		err = r.Commit(message, name, email)
		if err != nil {
			return err
		}

		fmt.Println("Changes committed successfully.")

		return nil
	},
}

func init() {
	commitCmd.Flags().StringVarP(&message, "message", "m", "Update configuration files", "The message when committing changes.")
	commitCmd.Flags().StringVarP(&name, "name", "n", "", "The author name")
	commitCmd.Flags().StringVarP(&email, "email", "e", "", "The author email")

	dotfCmd.AddCommand(commitCmd)
}
