package middleware

import (
	"fmt"

	"github.com/willyrgf/rerest/api/auth"
	log "github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
)

const (
	apiKeyArg = "api_key"
)

func getAPIKey(ctx *fasthttp.RequestCtx) (apiKey string) {
	args := ctx.QueryArgs()
	b := args.Peek(apiKeyArg)
	if b == nil {
		return
	}
	apiKey = string(b)
	log.Debugf("middleware.getAPIKey() apiKey=%+v", apiKey)
	return
}

// GetKeyOfCtx return the key of api key in db
func GetKeyOfCtx(ctx *fasthttp.RequestCtx) (key string) {
	apiKey := getAPIKey(ctx)
	key = fmt.Sprintf("%s:%s", "api_key", apiKey)
	return
}

// Auth middleware
func Auth(h fasthttp.RequestHandler) fasthttp.RequestHandler {
	return fasthttp.RequestHandler(func(ctx *fasthttp.RequestCtx) {
		apiKey := getAPIKey(ctx)
		log.Debugf("middleware.Auth(): apiKey=%+v", apiKey)
		if auth.IsAuthorized(apiKey) {
			log.Debugf("middleware.Auth(): authorized apiKey=%+v", apiKey)
			h(ctx)
			return
		}
		ctx.Error(fasthttp.StatusMessage(fasthttp.StatusUnauthorized), fasthttp.StatusUnauthorized)
	})
}
