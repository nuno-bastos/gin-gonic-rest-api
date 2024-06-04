//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"

	server "golang-gin-api/api"
	controller "golang-gin-api/api/controller"
	config "golang-gin-api/config"
	db "golang-gin-api/db"
	repository "golang-gin-api/repo"
	service "golang-gin-api/service"
)

func InitializeAPI(cfg config.Config) (*server.ServerHTTP, error) {
	wire.Build(db.ConnectDatabase,
		repository.NewTagRepository,
		service.NewTagService,
		controller.NewTagController,
		server.NewServerHTTP)

	return &server.ServerHTTP{}, nil
}
