package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "IP tracker version",
		Long: `The current version of the IP Tracker`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("v1.0.1")
		},
	}
)

// Execute executes the root command.
func init() {
	rootCmd.AddCommand(versionCmd)
}