package main

import (
	"api-service-test/internal/config"
	"api-service-test/internal/handlers"
	"api-service-test/internal/repository"
	"api-service-test/internal/service"
	"api-service-test/migrations"
	"fmt"
	"log"
	"net/http"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(fmt.Sprintf("failed to load config: %v", err))
	}

	db, err := gorm.Open(postgres.Open(cfg.Postgres.DSN), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("failed to connect database: %v", err))
	}

	if err = migrations.RunMigrations(cfg.Postgres.DSN); err != nil {
		panic(fmt.Sprintf("failed to run migrations: %v", err))
	}

	repo := repository.NewRepository(db)
	svc := service.NewService(repo)
	router := handlers.NewRouter(svc)

	srv := &http.Server{
		Addr:    cfg.Server.Host + ":" + cfg.Server.Port,
		Handler: router,
	}

	if err = srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}
