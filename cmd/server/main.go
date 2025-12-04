package main

import (
	"api-service-test/internal/config"
	"api-service-test/internal/handlers"
	"api-service-test/internal/repository"
	"api-service-test/internal/service"
	"api-service-test/migrations"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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

	if err := migrations.RunMigrations(cfg.Postgres.DSN); err != nil {
		panic(fmt.Sprintf("failed to run migrations: %v", err))
	}

	repo := repository.NewRepository(db)
	svc := service.NewService(repo)
	router := handlers.NewRouter(svc)

	srv := &http.Server{
		Addr:    cfg.Server.Host + ":" + cfg.Server.Port,
		Handler: router,
	}

	go func() {
		fmt.Printf("Starting server on %s\n", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	fmt.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		panic(fmt.Sprintf("server forced to shutdown: %v", err))
	}

	fmt.Println("Server exited gracefully")
}
