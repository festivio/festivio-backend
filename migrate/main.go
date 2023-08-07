package main

import (
	"github.com/festivio/festivio-backend/config"
	"github.com/festivio/festivio-backend/internal/database"
)

func main() {
	cfg := config.MustLoad()

	db, err := database.NewPsqlDB(cfg)
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate()
	if err != nil {
		println("Migration failed")
		panic(err)
	}
	println("Migration complete")
}
