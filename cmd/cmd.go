package cmd

import "github.com/spf13/cobra"

type Command = cobra.Command

var (
	rootCmd = Command{
		Use:   "stock-cli",
		Short: "An app for stock do some funny things.",
		Long:  "An app for stock that sync some stock info and note.",
	}
)

func init() {

}
