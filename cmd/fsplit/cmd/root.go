package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/shohi/glory/cmd/fsplit/pkg/actor"
	"github.com/shohi/glory/cmd/fsplit/pkg/config"
	"github.com/shohi/glory/cmd/fsplit/pkg/merge"
	"github.com/shohi/glory/cmd/fsplit/pkg/split"
	"github.com/spf13/cobra"
)

var conf = config.Config{}

var rootCmd = &cobra.Command{
	Use:   "fsplit",
	Short: "fsplit splits file into even parts",
	Run: func(cmd *cobra.Command, args []string) {
		var ac actor.Actor
		if conf.IsMerge {
			ac = merge.New(conf, args)
		} else {
			ac = split.New(conf, args)
		}

		err := ac.Act()
		if err != nil {
			log.Printf("err: %v", err)
			os.Exit(-1)
		}
	},
}

// setupFlags sets flags for comand line
func setupFlags(cmd *cobra.Command) {
	flagSet := cmd.Flags()

	// Server configuration
	flagSet.BoolVarP(&conf.IsMerge, "merge", "m", false, "whether to enable merge mode. merge given input list")
	flagSet.IntVarP(&conf.Number, "num", "n", 2, "split number")
	flagSet.StringVarP(&conf.Pattern, "pattern", "p", "", "file pattern for merge under current directory. If set, arguments are ignored")
}

// Execute is the entrance.
func Execute() {
	setupFlags(rootCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
