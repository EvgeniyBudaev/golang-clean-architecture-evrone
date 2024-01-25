package useCase

import (
	"context"

	"github.com/EvgeniyBudaev/golang-clean-architecture-evrone/internal/entity/room"
	"go.uber.org/zap"
)

func (u *UseCase) ListRooms(ctx context.Context) ([]*room.Room, error) {
	rooms, err := u.roomService.Fetch(ctx)
	if err != nil {
		u.logger.Debug("error func ListRooms, method Fetch by path internal/useCase/room/room.go", zap.Error(err))
		return nil, err
	}

	return rooms, nil
}
