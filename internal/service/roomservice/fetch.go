package roomservice

import (
	"context"

	"github.com/EvgeniyBudaev/golang-clean-architecture-evrone/internal/entity/room"
)

func (s *Service) Fetch(ctx context.Context) ([]*room.Room, error) {
	return s.roomRepo.Fetch(ctx)
}
