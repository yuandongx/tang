package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	var force bool
	var _init bool
	initDb := &Command{
		Use:  "db",
		Long: "init database and create tables that app needs",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("init db...")
			fmt.Println(args, force, _init)
		},
	}
	initDb.Flags().BoolVarP(&force, "force", "f", false, "force create or delete.")
	initDb.Flags().BoolVarP(&_init, "init", "i", false, "init a database tables.")
	rootCmd.AddCommand(initDb)
}
