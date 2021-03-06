// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package wire

import (
	"github.com/trewanek/go-with-ddd-example/application/adapter"
	"github.com/trewanek/go-with-ddd-example/application/usecase"
	"github.com/trewanek/go-with-ddd-example/domain/service"
	"github.com/trewanek/go-with-ddd-example/infrastructure/factory/memory"
	memory2 "github.com/trewanek/go-with-ddd-example/infrastructure/persistence/memory"
)

// Injectors from wire.go:

func InitializeUserRegisterUseCase() *usecase.RegisterUserUseCase {
	iUserFactory := memory.NewInMemoryUserFactory()
	iUserRepository := memory2.NewInMemoryUserRepository()
	iUserService := service.NewUserService(iUserRepository)
	userRegisterUseCase := usecase.NewUserRegisterUseCase(iUserFactory, iUserService, iUserRepository)
	return userRegisterUseCase
}

func InitializeUserService() adapter.IUserService {
	iUserRepository := memory2.NewInMemoryUserRepository()
	iUserService := service.NewUserService(iUserRepository)
	return iUserService
}
