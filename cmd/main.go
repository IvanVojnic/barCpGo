package main

import (
	"barCpGo/pkg/handler"
	"context"
	"github.com/segmentio/kafka-go"
	"github.com/spf13/viper"
	"log"
	"time"
)

func main() {
	conn, _ := kafka.DialLeader(context.Background(), "tcp", "localhost", "quickstart-events", 0)
	conn.SetWriteDeadline(time.Now().Add(time.Second * 20))
	conn.WriteMessages(kafka.Message{Value: []byte("hi")})
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		log.Fatalf("failed to init db", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(cpGo.Server)
	if err := srv.Run("3000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error", err.Error())
	}
}
