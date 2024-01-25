package useCase

import (
	"context"

	"github.com/EvgeniyBudaev/golang-clean-architecture-evrone/internal/entity/room"
	"go.uber.org/zap"
)

type createRoomRequest interface {
	GetRoomName() string
	GetTitle() string
}

func (u *UseCase) CreateRoom(ctx context.Context, request interface{ createRoomRequest }) (*room.Room, error) {
	room := room.NewRoom(request.GetRoomName(), request.GetTitle())

	room, err := u.roomService.Create(ctx, room)
	if err != nil {
		u.logger.Debug("error func CreateRoom, method Create by path internal/useCase/room/room.go", zap.Error(err))
		return nil, err
	}

	return room, nil
}
