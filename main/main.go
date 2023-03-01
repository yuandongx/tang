package main

import (
	"strings"
	"tang/api"
	"tang/scheduler"

	"github.com/gin-gonic/gin"
)

func SetupApp() {
	handlers := api.LoadApi()
	app := gin.New()
	apiv1 := app.Group("apiV1")
	for _, h := range handlers {
		method := strings.ToLower(h.Method)
		switch method {
		case "get":
			apiv1.GET(h.Path, h.Handler)
		case "post":
			apiv1.POST(h.Path, h.Handler)
		case "delete":
			apiv1.DELETE(h.Path, h.Handler)
		case "patch":
			apiv1.PATCH(h.Path, h.Handler)
		}
	}
	app.Run(":8080")
}

func main() {
	go scheduler.Run()
	SetupApp()
}
