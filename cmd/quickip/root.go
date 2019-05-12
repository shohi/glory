package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "quickip",
	Short: "quickip finds ip for given domain with least latency",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: add implementation
		fmt.Println("Hello QuickIP")
	},
}

// Execute is the entrance.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
