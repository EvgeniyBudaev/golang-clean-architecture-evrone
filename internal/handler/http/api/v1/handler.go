package v1

import (
	"context"

	"github.com/EvgeniyBudaev/golang-clean-architecture-evrone/internal/entity/room"
	"github.com/EvgeniyBudaev/golang-clean-architecture-evrone/internal/logger"
)

type UseCase interface {
	CreateRoom(ctx context.Context, request interface{ createRoomRequest }) (*room.Room, error)
	ListRooms(ctx context.Context) ([]*room.Room, error)
}

type Handler struct {
	logger  logger.Logger
	useCase UseCase
}

func NewHandler(logger logger.Logger, useCase UseCase) *Handler {
	return &Handler{
		logger:  logger,
		useCase: useCase,
	}
}
