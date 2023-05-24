package handlers

import (
	"tang/sqldb"

	"github.com/gin-gonic/gin"
)

var db *sqldb.Session

const (
	GET = iota
	POST
	DELETE
	PATCH
)

type GContext = gin.Context
type HandlerFunc = gin.HandlerFunc

type Handler struct {
	Path   string
	Method int
	HandlerFunc
}

func LoadHandlers(sqldb *sqldb.Session) []Handler {
	db = sqldb
	handlers := []Handler{}
	return handlers
}
