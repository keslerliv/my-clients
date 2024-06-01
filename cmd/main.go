package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/keslerliv/my-clients/config"
	"github.com/keslerliv/my-clients/internal/crons"
	"github.com/keslerliv/my-clients/internal/routes"
	"github.com/keslerliv/my-clients/pkg/db"
)

// flag for dev mode
var (
	DevMode = flag.Bool("dev", false, "dev mode")
)

func main() {
	flag.Parse()
	config.InitConfig(*DevMode)

	r := routes.LoadRoutes()

	// create database migrations
	db.MakeMigrations()

	// load crons
	crons.LoadCrons()

	http.ListenAndServe(fmt.Sprintf(":%s", config.Config.App.Port), r)
}
