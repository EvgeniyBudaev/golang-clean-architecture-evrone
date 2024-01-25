package roomrepo

import (
	"context"

	"github.com/EvgeniyBudaev/golang-clean-architecture-evrone/internal/entity/room"
	"go.uber.org/zap"
)

func (r *Repository) Create(ctx context.Context, room *room.Room) (*room.Room, error) {
	query := "INSERT INTO rooms (room_name, title) VALUES ($1, $2) RETURNING id"

	err := r.conn.QueryRowContext(ctx, query, room.RoomName, room.Title).Scan(&room.ID)
	if err != nil {
		r.logger.Debug("error func CreateRoom, method QueryRowContext by path internal/db/room/room.go",
			zap.Error(err))
		return nil, err
	}

	return room, nil
}
