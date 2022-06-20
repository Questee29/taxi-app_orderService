package main

import (
	"errors"
	"log"

	config "github.com/Questee29/taxi-app_orderService/configs"
	"github.com/Questee29/taxi-app_orderService/database"

	_ "github.com/Questee29/taxi-app_orderService/migrations"
	server "github.com/Questee29/taxi-app_orderService/pkg/grpcServer"
	handlers "github.com/Questee29/taxi-app_orderService/pkg/grpcServer/handler"
	"github.com/Questee29/taxi-app_orderService/pkg/repository"
	service "github.com/Questee29/taxi-app_orderService/pkg/service"
)

func main() {
	config, err := config.LoadConfig("app", ".")
	if err != nil {
		log.Fatal("cannot load config", err)
	}

	db, err := database.New()
	if err != nil {
		log.Fatalln(errors.New(`failed to load database`))
	}
	repository := repository.New(db)
	service := service.New(repository)
	grpcOrderHandler := handlers.NewOrderHandler(service)
	grpcServ := server.NewServer(server.Deps{
		OrderHandler: grpcOrderHandler,
	})
	if err := grpcServ.ListenAndServe(config.Server.Port); err != nil {
		log.Printf("grpc ListenAndServe err %s", err)
	}

}
