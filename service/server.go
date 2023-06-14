package service

import (
	"tang/service/view"
	"tang/sqldb"
)

func LoadApp(db *sqldb.Session) *view.Echo {
	hds := view.LoadHandlers()
	app := view.New()
	apiv1 := app.Group("/apiV1")
	{
		for _, h := range hds {
			switch h.Htype {
			case view.GET:
				apiv1.GET(h.Path, h.Function)
			case view.POST:
				apiv1.POST(h.Path, h.Function)
			case view.DELETE:
				apiv1.DELETE(h.Path, h.Function)
			case view.PATCH:
				apiv1.PATCH(h.Path, h.Function)
			default:
				apiv1.GET(h.Path, h.Function)
			}
		}
	}
	return app
}
