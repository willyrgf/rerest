package handler

import (
	"github.com/willyrgf/rerest/api/controllers"
	"github.com/willyrgf/rerest/api/middleware"
	"github.com/willyrgf/rerest/config/router"
)

// MakeHandler set all handlers to api
func MakeHandler() {
	r := router.Get()
	r.GET("/:key_prefix/:key_suffix", middleware.Auth(controllers.GetKey))
}
