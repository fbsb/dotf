package cmd

import (
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

var dotfCmd = &cobra.Command{
	Use:   "dotf",
	Short: "Dotf is a tool for managing and syncing configuration files.",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		absPath, err := filepath.Abs(rootPath)
		if err != nil {
			return err
		}
		rootPath = absPath
		return nil
	},
}

var (
	rootPath string
)

func defaultRootPath() string {
	home, err := homedir.Dir()
	if err != nil {
		panic(err)
	}

	return home
}

func init() {
	dotfCmd.PersistentFlags().StringVarP(&rootPath, "root", "r", defaultRootPath(), "The path containing the .dotfiles directory")
}

func Execute() error {
	return dotfCmd.Execute()
}
