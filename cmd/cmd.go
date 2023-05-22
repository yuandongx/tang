package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

type Command = cobra.Command

var (
	rootCmd = Command{
		Use:   "stock",
		Short: "An app for stock do some funny things.",
		Long:  "An app for stock that sync some stock info and note.",
	}
)

func RunCmd() {
	err := rootCmd.Execute()
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
