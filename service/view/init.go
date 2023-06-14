package view

import (
	"tang/sqldb"

	"github.com/labstack/echo/v4"
)

var db *sqldb.Session

const (
	GET = iota
	POST
	DELETE
	PATCH
)

type EContext = echo.Context
type Echo = echo.Echo
type Function = echo.HandlerFunc
type Handler struct {
	Htype int
	Function
	Path string
}

func LoadHandlers() (handlers []Handler) {
	return
}

func New() *Echo {
	return echo.New()
}
