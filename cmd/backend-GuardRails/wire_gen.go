// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"backend/internal/biz"
	"backend/internal/conf"
	"backend/internal/data"
	"backend/internal/server"
	"backend/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	db, err := data.NewPostgresClient(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	dataData, err := data.NewData(db)
	if err != nil {
		return nil, nil, err
	}
	repositoryRepo := data.NewRepositoryRepo(dataData, logger)
	repositoryUsecase := biz.NewRepositoryUsecase(repositoryRepo, logger)
	repositoryService := service.NewRepositoryService(repositoryUsecase)
	scanRepo := data.NewScanRepo(dataData, logger)
	gitCloneKafkaRepo := data.NewGitCloneKafkaRepo(confData, logger)
	scanUsecase := biz.NewScanUsecase(scanRepo, gitCloneKafkaRepo, logger)
	fileContentKafkaRepo := data.NewFileContentKafkaRepo(confData, logger)
	gitCloneUsecase := biz.NewGitCloneUseCase(repositoryRepo, scanRepo, gitCloneKafkaRepo, fileContentKafkaRepo, logger)
	fileContentUsecase := biz.NewFileContentUseCase(scanRepo, fileContentKafkaRepo, logger)
	scanService := service.NewScanService(scanUsecase, gitCloneUsecase, fileContentUsecase)
	grpcServer := server.NewGRPCServer(confServer, repositoryService, scanService, logger)
	httpServer := server.NewHTTPServer(confServer, repositoryService, scanService, logger)
	app := newApp(logger, grpcServer, httpServer)
	return app, func() {
	}, nil
}
