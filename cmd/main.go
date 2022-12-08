package main

import (
	"barCpGo"
	"barCpGo/pkg/handler"
	"barCpGo/pkg/repository"
	"barCpGo/pkg/service"
	"bytes"
	"context"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/segmentio/kafka-go"
	"github.com/spf13/viper"
	"log"
	"strconv"
	"time"
)

func main() {
	conn, _ := kafka.DialLeader(context.Background(), "tcp", "localhost", "placeConn", 0)
	conn.SetReadDeadline(time.Now().Add(time.Second * 20))
	message, _ := conn.ReadMessage(1e6)
	fmt.Println(string(message.Value))
	id_place := bytes.NewBuffer(message.Value).String()
	var id_placeInt int
	id_placeInt, _ = strconv.Atoi(id_place)
	name, err :=
	if err != nil {
		return
	}
	fmt.Println(name)

	if err := initConfig(); err != nil {
		log.Fatalf("error init config", err.Error())
	}
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

	srv := new(barCpGo.Server)
	if err := srv.Run("5000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

/*
  migrate -path ./schema -database 'postgres://root:vojnic@localhost:5438/postgres?sslmode=disable' up
  docker run --name=bar-db -e POSTGRES_PASSWORD='vojnic' -p 5438:5432 -d --rm postgres
*/
