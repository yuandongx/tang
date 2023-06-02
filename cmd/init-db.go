package cmd

import (
	"fmt"
	model "tang/service/model"
	"tang/sqldb"

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
			initdb()
		},
	}
	initDb.Flags().BoolVarP(&force, "force", "f", false, "force create or delete.")
	initDb.Flags().BoolVarP(&_init, "init", "i", false, "init a database tables.")
	rootCmd.AddCommand(initDb)
}

func initdb() {
	session := sqldb.CreatePostgresSession("")
	m := model.XStock{}
	session.Create(m)
}
