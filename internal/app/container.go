package app

import (
	"database/sql"

	"github.com/EvgeniyBudaev/golang-clean-architecture-evrone/internal/adapter/psqlrepo/roomrepo"
	"github.com/EvgeniyBudaev/golang-clean-architecture-evrone/internal/logger"
	"github.com/EvgeniyBudaev/golang-clean-architecture-evrone/internal/service/roomservice"
	"github.com/EvgeniyBudaev/golang-clean-architecture-evrone/internal/useCase"
)

type Container struct {
	logger   logger.Logger
	postgres *sql.DB
}

func NewContainer(logger logger.Logger, mysql *sql.DB) *Container {
	return &Container{
		logger:   logger,
		postgres: mysql,
	}
}

func (c *Container) GetUseCase() *useCase.UseCase {
	return useCase.NewUseCase(c.logger, c.getRoomService())
}

func (c *Container) getPostgres() *sql.DB {
	return c.postgres
}

func (c *Container) getRoomRepo() *roomrepo.Repository {
	return roomrepo.NewRepository(c.logger, c.getPostgres())
}

func (c *Container) getRoomService() *roomservice.Service {
	return roomservice.NewService(c.getRoomRepo())
}
