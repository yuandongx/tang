package main

import (
	"tang/api"

	"github.com/gin-gonic/gin"
)

func SetupApp() {
	handlers := api.LoadApi()
	app := gin.New()
	apiv1 := app.Group("apiV1")
	for _, h := range handlers {
		switch h.Method {
		case "get", "GET", "Get":
			apiv1.GET(h.Path, *h.Handler)
		case "post", "POST", "Post":
			apiv1.POST(h.Path, *h.Handler)
		case "delete", "DELETE", "Delete":
			apiv1.DELETE(h.Path, *h.Handler)
		case "patch", "PATCH", "Patch":
			apiv1.PATCH(h.Path, *h.Handler)
		}
	}
	app.Run(":8080")
}

func main() {
	SetupApp()
}
