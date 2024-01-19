package main

import (
	"effective_mobile_junior/external/agify"
	"effective_mobile_junior/external/genderize"
	"effective_mobile_junior/external/logger"
	"effective_mobile_junior/external/nationalize"
	"effective_mobile_junior/external/postgres"
	"effective_mobile_junior/internal/handler"
	"effective_mobile_junior/internal/repository"
	"effective_mobile_junior/internal/service"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("File .env not found")
	}
}

func main() {
	client := http.Client{}

	// Берем уровень логгирования из .env. Если его нет - уровень info
	logLevel, exist := os.LookupEnv("LOG_LEVEL")
	if !exist {
		logLevel = "info"
	}
	l, err := logger.New(logLevel)
	if err != nil {
		l.Error("error read logging level",
			zap.String("given level", "test"),
			zap.String("error", err.Error()),
		)
	}

	l.Info("level logging", zap.String("level", l.Level().String()))

	// Подключение к БД
	dbConf := postgres.NewConfig()
	db, err := postgres.NewConnPool(dbConf)
	if err != nil {
		l.Error("Ошибка подключения к базе данных",
			zap.String("error", err.Error()),
		)
	}

	// Инициализация сторонних API сервисов
	agifyEngine := agify.NewEngine(agify.WithCustomClient(&client))
	genderizeEngine := genderize.NewEngine(genderize.WithCustomClient(&client))
	nationalizeEngine := nationalize.NewEngine(nationalize.WithCustomClient(&client))

	repo := repository.New(l, db)
	services := service.New(l, agifyEngine, nationalizeEngine, genderizeEngine, repo)
	handlers := handler.New(l, services)

	// Берем порт сервера из .env. Если его нет - уровень 8080
	addr, exist := os.LookupEnv("SERVER_PORT")
	if !exist {
		addr = "8080"
	}

	go func() {
		if err := http.ListenAndServe(":"+addr, handlers.Init()); err != nil {
			l.Fatal("errors serve handler",
				zap.String("port", addr),
				zap.String("error", err.Error()),
			)
		}
	}()

	l.Info("server start")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	l.Info("server shutdown")
}
