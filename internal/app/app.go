package app

import (
	"main.go/internal/handler"
	"main.go/internal/repository"
	"main.go/internal/server"
	"main.go/internal/service"
)

func Run() {

	//handler
	repository := repository.NewRepository()
	service := service.NewService(repository)
	handler := handler.NewHandler(service)

	//server
	server := server.NewServer(handler)
	err := server.Start()
	if err != nil {
		panic(err)
	}
}
