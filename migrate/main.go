package main

import (
	"github.com/festivio/festivio-backend/config"
	"github.com/festivio/festivio-backend/domain"
	"github.com/festivio/festivio-backend/internal/database"
)

func main() {
	cfg := config.MustLoad()

	db, err := database.NewPsqlDB(cfg)
	if err != nil {
		panic(err)
	}

	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	err = db.AutoMigrate(&domain.User{})
	if err != nil {
		println("Migration failed")
		panic(err)
	}
	println("Migration complete")
}
