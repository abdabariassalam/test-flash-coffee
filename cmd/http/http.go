package http

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bariasabda/test-flash-coffee/config"
	"github.com/bariasabda/test-flash-coffee/domain/repository"
	"github.com/bariasabda/test-flash-coffee/domain/service"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func Execute() {
	// Configuration
	var cfg config.Config

	err := cleanenv.ReadConfig("./config/config.yml", &cfg)
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// connect postgresql
	PostgresqlUrl := os.Getenv("PostgresqlUrl")
	if PostgresqlUrl == "" {
		PostgresqlUrl = cfg.Postgresql.Url
	}
	db, err := gorm.Open("postgres",PostgresqlUrl)
	if err != nil {
		log.Fatalf("Tidak Konek DB Errornya : %s", err)
	}

	repository := repository.New(db)
	service := service.New(repository)

	// Run
	handler := NewHandler(service)
	r := NewRouter(*handler)
	r.routes()
	port := ":" + os.Getenv("PORT")
	if port == ":" {
		port = cfg.Base.Port
	}
	r.router.Run(port)

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Fatalf("app - Run - signal: %s", s.String())
	default:
	}
}
