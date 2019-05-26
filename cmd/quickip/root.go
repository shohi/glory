package main

import (
	"os"
	"time"

	"github.com/spf13/cobra"

	log "github.com/sirupsen/logrus"

	"github.com/shohi/glory/cmd/quickip/pkg/config"
	"github.com/shohi/glory/cmd/quickip/pkg/search"
)

var conf = config.Config{}

var rootCmd = &cobra.Command{
	Use:   "quickip",
	Short: "quickip finds ip for given domain with least latency",
	Run: func(cmd *cobra.Command, args []string) {
		initLog(conf.LogLevel)

		if len(args) == 0 {
			log.Error("arguments must not empty")
			os.Exit(-1)
		}

		startT := time.Now()
		defer func() {
			log.WithField("duration", time.Since(startT)).Info()
		}()

		// TODO: Add mutliple args support
		search.Search(conf, args)
	},
}

// setupFlags sets flags for comand line
func setupFlags(cmd *cobra.Command) {
	flagSet := cmd.Flags()

	// Server configuration
	flagSet.BoolVarP(&conf.ShowLocation, "location", "l", false, "whether to show ip location")
	flagSet.StringVar(&conf.LogLevel, "loglevel", "INFO", "log level")
}

// Execute is the entrance.
func Execute() {
	setupFlags(rootCmd)

	_ = rootCmd.Execute()
}

func initLog(level string) {
	l, err := log.ParseLevel(level)
	if err != nil {
		l = log.InfoLevel
	}

	log.SetLevel(l)
	log.SetFormatter(&log.TextFormatter{
		DisableColors:    true,
		DisableTimestamp: true,
		// FullTimestamp:    true,
	})
}
