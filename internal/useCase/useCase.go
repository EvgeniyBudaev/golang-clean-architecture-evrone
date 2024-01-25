package useCase

import (
	"context"

	"github.com/EvgeniyBudaev/golang-clean-architecture-evrone/internal/entity/room"
	"github.com/EvgeniyBudaev/golang-clean-architecture-evrone/internal/logger"
)

type roomService interface {
	Create(ctx context.Context, room *room.Room) (*room.Room, error)
	Fetch(ctx context.Context) ([]*room.Room, error)
}

type UseCase struct {
	logger      logger.Logger
	roomService roomService
}

func NewUseCase(logger logger.Logger, roomService roomService) *UseCase {
	return &UseCase{
		logger:      logger,
		roomService: roomService,
	}
}
