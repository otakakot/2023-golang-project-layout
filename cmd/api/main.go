package main

import (
	"os"

	"github.com/otakakot/2023-golang-project-layout/internal/adapter/controller"
	"github.com/otakakot/2023-golang-project-layout/internal/adapter/gateway"
	"github.com/otakakot/2023-golang-project-layout/internal/application/interactor"
	"github.com/otakakot/2023-golang-project-layout/internal/driver/postgres"
	"github.com/otakakot/2023-golang-project-layout/internal/driver/server"
	"github.com/otakakot/2023-golang-project-layout/pkg/api"
)

func main() {
	db, err := postgres.New(os.Getenv("POSTGRES_URL"))
	if err != nil {
		panic(err)
	}

	gw := gateway.NewTodo(db)

	uc := interactor.NewTodo(gw)

	ctl := controller.NewTodo(uc)

	hdl, err := api.NewServer(ctl)
	if err != nil {
		panic(err)
	}

	srv := server.NewServer("8080", hdl)

	srv.Run()
}
