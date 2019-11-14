package controllers

import (
	"fmt"

	"github.com/willyrgf/rerest/api/middleware"
	"github.com/willyrgf/rerest/api/responses"
	"github.com/willyrgf/rerest/config"
	"github.com/willyrgf/rerest/pkg/account"
	"github.com/willyrgf/rerest/pkg/repository"
	log "github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
)

func getKeyPartsFromURL(ctx *fasthttp.RequestCtx) (kp, ks, k string, ok bool) {
	var v interface{}

	v = ctx.UserValue("key_prefix")
	kp, ok = v.(string)
	if !ok {
		return
	}

	v = ctx.UserValue("key_suffix")
	ks, ok = v.(string)
	if !ok {
		return
	}

	k = fmt.Sprintf("%s:%s", kp, ks)
	if k != "" || kp != "" || ks != "" {
		ok = true
	}

	return
}

// GetKey is a controller for get any key
func GetKey(ctx *fasthttp.RequestCtx) {
	keyPrefix, keySuffix, key, ok := getKeyPartsFromURL(ctx)
	if !ok {
		ctx.Error("Unsupported path", fasthttp.StatusNotFound)
		return
	}

	keyConf, ok := config.RerestConf.Access[keyPrefix]
	if !ok {
		ctx.Error("Unsupported path", fasthttp.StatusNotFound)
		return
	}

	log.Debugf("controllers.GetKey(): key_prefix=%s; key_suffix=%s; key=%s keyconf=%+v\n", keyPrefix, keySuffix, key, keyConf)

	switch keyConf.Type {
	case "set":
		members, err := repository.GetFromSet(key)
		if err != nil {
			ctx.Error("error on get key", fasthttp.StatusInternalServerError)
		}
		log.Debugf("controllers.GetKey(): set; members=%+v", members)

		answered := responses.Strings(ctx, members)
		if answered {
			go account.Response(middleware.GetKeyOfCtx(ctx))
		}
		return

	default:
		ctx.Error("Type of key unsupported", fasthttp.StatusInternalServerError)
		return
	}
}
