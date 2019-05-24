package main

import (
	"fmt"
	"log"
	"os"

	"github.com/shohi/glory/cmd/fsplit/pkg/actor"
	"github.com/shohi/glory/cmd/fsplit/pkg/combine"
	"github.com/shohi/glory/cmd/fsplit/pkg/config"
	"github.com/shohi/glory/cmd/fsplit/pkg/split"
	"github.com/spf13/cobra"
)

var conf = config.Config{}

var rootCmd = &cobra.Command{
	Use:   "fsplit",
	Short: "fsplit splits file into equal parts",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Println("arguments must not be empty")
			os.Exit(1)
		}

		var ac actor.Actor
		if conf.IsCombine {
			ac = combine.New(conf, args)
		} else {
			ac = split.New(conf, args)
		}

		err := ac.Act()
		if err != nil {
			log.Printf("err: %v", err)
		}
	},
}

// setupFlags sets flags for comand line
func setupFlags(cmd *cobra.Command) {
	flagSet := cmd.Flags()

	// Server configuration
	flagSet.BoolVarP(&conf.IsCombine, "combine", "c", false, "combine given input list")
	flagSet.IntVarP(&conf.Number, "num", "n", 2, "split number")
}

// Execute is the entrance.
func Execute() {
	setupFlags(rootCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
