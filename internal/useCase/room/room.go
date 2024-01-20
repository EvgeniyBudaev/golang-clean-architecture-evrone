package room

import (
	"context"
	"github.com/EvgeniyBudaev/golang-clean-architecture-evrone/internal/db/room"
	roomEntity "github.com/EvgeniyBudaev/golang-clean-architecture-evrone/internal/entity/room"
	"github.com/EvgeniyBudaev/golang-clean-architecture-evrone/internal/logger"
	"go.uber.org/zap"
)

type UseCaseRoom struct {
	db *room.PGRoomDB
}

func NewUseCaseRoom(db *room.PGRoomDB) *UseCaseRoom {
	return &UseCaseRoom{
		db: db,
	}
}

type CreateRoomRequest struct {
	RoomName string `json:"roomName"`
	Title    string `json:"title"`
}

func (uc *UseCaseRoom) CreateRoom(ctx context.Context, r CreateRoomRequest) (*roomEntity.Room, error) {
	roomRequest := &roomEntity.Room{
		RoomName: r.RoomName,
		Title:    r.Title,
	}
	newRoom, err := uc.db.CreateRoom(ctx, roomRequest)
	if err != nil {
		logger.Log.Debug("error func CreateRoom, method Create by path internal/useCase/room/room.go",
			zap.Error(err))
		return nil, err
	}
	return newRoom, nil
}

func (uc *UseCaseRoom) GetRoomList(ctx context.Context) ([]*roomEntity.Room, error) {
	response, err := uc.db.SelectRoomList(ctx)
	if err != nil {
		logger.Log.Debug("error func GetRoomList, method SelectRoomList by path internal/useCase/room/room.go",
			zap.Error(err))
		return nil, err
	}
	return response, nil
}
