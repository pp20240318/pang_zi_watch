package main

import (
	"log"

	"watch/backend/internal/config"
	"watch/backend/internal/db"
	"watch/backend/internal/router"
	"watch/backend/internal/service"
)

func main() {
	cfg := config.Load()
	conn := db.Connect(cfg)
	service.EnsureDefaultAdmin(conn)
	service.SeedCatalog(conn)
	service.SeedContentPages(conn)

	r := router.Setup(conn, cfg)
	log.Printf("watch admin API listening on %s", cfg.Addr())
	if err := r.Run(cfg.Addr()); err != nil {
		log.Fatal(err)
	}
}
