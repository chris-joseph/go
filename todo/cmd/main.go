package main

import (
	"todo/internal/api"
	"todo/pkg/config"
	"todo/pkg/data"
)

func main() {

	cfg := config.New()
	db := data.NewMongoconnection(cfg)
	defer db.Disconnect()
	application := api.New(cfg, db.Client)
	application.Start()
}
