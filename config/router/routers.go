package router

import (
	"github.com/buaazp/fasthttprouter"
)

var router *fasthttprouter.Router

func initRouter() {
	router = fasthttprouter.New()
}

// Get router for ReREST
func Get() *fasthttprouter.Router {
	if router == nil {
		initRouter()
	}
	return router
}
