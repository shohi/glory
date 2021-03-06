package cmd

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/shohi/glory/pkg/httputil"
	"github.com/shohi/glory/pkg/serde"
	"github.com/spf13/cobra"
)

type Config struct {
	URL   string
	Send  bool
	Quiet bool
}

var conf Config
var client http.Client

var rootCmd = &cobra.Command{
	Use:   "resp",
	Short: "redis RESP cli",
	RunE:  run,
}

// setupFlags sets flags for comand line
func setupFlags(cmd *cobra.Command) {
	flagSet := cmd.Flags()

	// Server configuration
	flagSet.BoolVarP(&conf.Send, "send", "s", true, "send redis commands to url")
	flagSet.StringVarP(&conf.URL, "url", "u", "", "redis compatible http server address")
	flagSet.BoolVarP(&conf.Quiet, "quiet", "q", true, "quiet mode which will suppress all outputs")
}

// Execute is the entrance.
func Execute() {
	setupFlags(rootCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func send(conf Config, args []string) error {
	if len(args) == 0 {
		return errors.New("No command")
	}
	var cName string
	var cArgs []string

	if len(args) > 1 {
		cName, cArgs = args[0], args[1:]
	} else {
		cName = args[0]
	}

	data := serde.SerializeRawRESP(cName, cArgs)
	req, err := http.NewRequest("POST", conf.URL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	if !conf.Quiet {
		fmt.Println(resp.StatusCode)
	}

	return httputil.DiscardBodyAndClose(resp)
}

func get(conf Config) error {
	// TODO
	log.Printf("TODO")
	return nil
}

func run(cmd *cobra.Command, args []string) error {
	if conf.Send {
		return send(conf, args)
	}

	return get(conf)
}
