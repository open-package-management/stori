package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	// Version is the semantic version of the binary.
	Version = "undefined"

	// Commit is the target commit SHA of the binary.
	Commit = "undefined"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version information for Stori.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("stori version: %s\n", Version)
		fmt.Printf("commit: %s\n", Commit)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
