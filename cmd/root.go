package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "cobra-cli",
		Short: "IP Tracker CLI Application",
		Long: `IP Tracker CLI application built from ground-up.`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}