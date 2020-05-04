package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/valyala/fasthttp"
	"github.com/willyrgf/rerest/api/handler"
	"github.com/willyrgf/rerest/config"
	"github.com/willyrgf/rerest/config/router"

	log "github.com/sirupsen/logrus"
)

func init() {
	dev := flag.Bool("dev", false, "Set the environment to dev.")
	trace := flag.Bool("trace", false, "Enable trace.")
	flag.Parse()

	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})

	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)

	if *dev {
		log.SetLevel(log.DebugLevel)
		log.Debug("init(): dev environment")
	}

	if *trace {
		log.SetReportCaller(true)
		log.Debug("init(): trace enabled")
	}
}

func main() {
	flag.Parse()
	config.Load()
	router := router.Get()
	handler.MakeHandler()

	log.Infof("main(): start the api; config.RerestConf=%+v", config.RerestConf)
	server := fmt.Sprintf("%s:%d", config.RerestConf.API.HTTPHost, config.RerestConf.API.HTTPPort)
	log.Fatal(fasthttp.ListenAndServe(server, router.Handler))
}
