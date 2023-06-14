package cmd

import "tang/service"

func RunServer() {
	app := service.LoadApp(nil)
	app.Start(":8080")
}
