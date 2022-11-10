package api

import (
	"tang/db"

	"github.com/gin-gonic/gin"
)

type Context = gin.Context
type HandlerFunc = gin.HandlerFunc
type Handler struct {
	Path    string
	Handler *HandlerFunc
	Method  string
}

var mgdb db.MongoClient

func LoadApp() []Handler {
	return []Handler{}
}

func init() {
	mgdb = db.MongoClient{Uri: "mongodb://root:example@106.75.63.248:27017/"}
}
