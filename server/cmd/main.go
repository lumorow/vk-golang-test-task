package main

import (
	"context"
	"errors"
	"filmlib/server/internal/handler"
	"filmlib/server/internal/repository"
	"filmlib/server/internal/service"
	"filmlib/server/pkg/server"
	"os"
	"os/signal"
	"syscall"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/github"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// @title Film library
// @version 1.0
// description Film library application management
// @host localhost:8000
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @Description JWT token (example: Bearer <your_jwt_token_here>)
// @TokenUrl http://localhost:8000/auth/sign-in

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := godotenv.Load("./server/configs/.env"); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}
	viper.AutomaticEnv()

	dbConn, err := repository.NewPostgresDB(repository.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	})
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}
	defer func() {
		if err := dbConn.Close(); err != nil {
			logrus.Errorf("error occured on db connection close: %s", err.Error())
		}
	}()

	if err = repository.Migrate(dbConn); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		logrus.Fatalf("error init sql tables: %s", err.Error())
	}

	authRepos := repository.NewAuthorizationRepository(dbConn)
	actorRepos := repository.NewActorRepository(dbConn)
	filmRepos := repository.NewFilmRepository(dbConn)

	authService := service.NewAuthorizationService(authRepos)
	actorService := service.NewActorService(actorRepos, filmRepos)
	filmService := service.NewFilmService(actorRepos, filmRepos)

	handlers := handler.NewHandler(authService, actorService, filmService)

	srv := new(server.Server)

	go func() {
		if err := srv.Run(handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	logrus.Print("Film library application Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("Film library application Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}
}
