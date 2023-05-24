package service

import (
	"tang/service/handlers"
	"tang/sqldb"

	"github.com/gin-gonic/gin"
)

func LoadApp(db *sqldb.Session) *gin.Engine {
	hds := handlers.LoadHandlers(db)
	app := gin.Default()
	apiv1 := app.Group("/apiV1")
	{
		for _, h := range hds {
			switch h.Method {
			case handlers.GET:
				apiv1.GET(h.Path, h.HandlerFunc)
			case handlers.POST:
				apiv1.POST(h.Path, h.HandlerFunc)
			case handlers.DELETE:
				apiv1.DELETE(h.Path, h.HandlerFunc)
			case handlers.PATCH:
				apiv1.PATCH(h.Path, h.HandlerFunc)
			default:
				apiv1.GET(h.Path, h.HandlerFunc)
			}
		}
	}
	return app
}
