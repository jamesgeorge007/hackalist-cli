package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var version, commitHash string

var versionCommand = &cobra.Command{
	Use: "version",
	Short: "Shows the current hackalist version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("version: %s\ncommit: %s\n", version, commitHash)
	},
}

func init() {
	rootCmd.AddCommand(versionCommand)
}
