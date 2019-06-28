package main

import (
	"fmt"
	"os"

	"github.com/shohi/glory/cmd/allup/pkg/config"
	"github.com/spf13/cobra"
)

var conf = config.Config{}

var rootCmd = &cobra.Command{
	Use:   "allup",
	Short: "update all repos under given folder",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO
	},
}

// setupFlags sets flags for comand line
func setupFlags(cmd *cobra.Command) {
	flagSet := cmd.Flags()

	// Server configuration
	flagSet.StringVarP(&conf.Dir, "dir", "d", ".", "repos base directory")
}

// Execute is the entrance.
func Execute() {
	setupFlags(rootCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
