package roomservice

import (
	"context"

	"github.com/EvgeniyBudaev/golang-clean-architecture-evrone/internal/entity/room"
)

type roomRepo interface {
	Create(ctx context.Context, room *room.Room) (*room.Room, error)
	Fetch(ctx context.Context) ([]*room.Room, error)
}

type Service struct {
	roomRepo roomRepo
}

func NewService(roomRepo roomRepo) *Service {
	return &Service{
		roomRepo: roomRepo,
	}
}
