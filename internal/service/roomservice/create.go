package roomservice

import (
	"context"

	"github.com/EvgeniyBudaev/golang-clean-architecture-evrone/internal/entity/room"
)

func (s *Service) Create(ctx context.Context, room *room.Room) (*room.Room, error) {
	return s.roomRepo.Create(ctx, room)
}
