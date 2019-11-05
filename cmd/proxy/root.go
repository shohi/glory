package main

import (
	"fmt"
	"os"

	"github.com/shohi/glory/pkg/proxy"
	"github.com/spf13/cobra"
)

// TODO: use more general configuration, e.g.
// [{
//     pattern: xxx
//     addr: xxx
// }]

var conf proxy.Config

var rootCmd = &cobra.Command{
	Use:   "proxy",
	Short: "proxy redirects all requests to backend server with configured rules",
	RunE:  runProxy,
}

// setupFlags sets flags for comand line
func setupFlags(cmd *cobra.Command) {
	fs := cmd.Flags()

	// Server configuration
	fs.IntVar(&conf.Port, "port", 9302, "proxy server listen port")
	fs.StringVar(&conf.Rules, "rules", "", `redirction rules. multiple rules are separated by comma, e.g. `)
}

// Execute is the entrance.
func Execute() {
	setupFlags(rootCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func runProxy(cmd *cobra.Command, args []string) error {
	srv, err := proxy.NewServer(conf)
	if err != nil {
		return err
	}
	return srv.Start()
}
