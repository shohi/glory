package main

import (
	"bytes"
	"errors"
	"fmt"
	"os"

	"github.com/bsm/redeo/resp"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "resp",
	Short: "redis RESP cli",
	RunE:  runSerde,
}

// Execute is the entrance.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func runSerde(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return errors.New("No command")
	}

	cName := args[0]
	var cArgs []string
	if len(args) > 1 {
		cArgs = args[1:]
	}

	var buf bytes.Buffer
	w := resp.NewRequestWriter(&buf)
	w.WriteCmdString(cName, cArgs...)
	_ = w.Flush()

	fmt.Printf("%q\n", buf.String())

	return nil
}
