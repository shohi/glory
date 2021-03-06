package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"

	log "github.com/sirupsen/logrus"

	"github.com/shohi/glory/cmd/quickip/pkg/config"
	"github.com/shohi/glory/cmd/quickip/pkg/search"
)

var conf = config.Config{}
var showVersion bool
var gitCommit = "not set"

var rootCmd = &cobra.Command{
	Use:   "quickip",
	Short: "quickip finds ip for given domain with least latency",
	Run: func(cmd *cobra.Command, args []string) {
		if showVersion {
			fmt.Printf("version: %v\n", gitCommit)
			os.Exit(0)
		}

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
	flagSet.BoolVar(&conf.ShowLocation, "location", false, "show ip location")
	flagSet.BoolVarP(&conf.ShowLatency, "latency", "l", false, "show average latency")
	flagSet.StringVar(&conf.LogLevel, "loglevel", "INFO", "log level")

	flagSet.IntVar(&conf.PingCount, "pingcount", 3, "ping count")

	flagSet.BoolVarP(&showVersion, "version", "v", false, "show version")
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
