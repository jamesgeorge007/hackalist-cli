package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var version = "0.0.1"

var versionCommand = &cobra.Command{
	Use: "version",
	Short: "Shows the current hackalist version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("version: %s\n", version)
	},
}

func init() {
	rootCmd.AddCommand(versionCommand)
}
