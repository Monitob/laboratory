// Package cmd handles the command-line interface for laboratory
package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(reactCmd)
}

// RootCmd is the root for all hello commands.
var RootCmd = &cobra.Command{
	Use:           "laboratory",
	Short:         "abBc",
	Long:          `Currently, just write your polymere.`,
	SilenceErrors: true,
}
