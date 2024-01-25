package roomrepo

import (
	"context"

	"github.com/EvgeniyBudaev/golang-clean-architecture-evrone/internal/entity/room"
	"go.uber.org/zap"
)

func (r *Repository) Fetch(ctx context.Context) ([]*room.Room, error) {
	query := "SELECT id, room_name, title FROM rooms"

	rows, err := r.conn.QueryContext(ctx, query)
	if err != nil {
		r.logger.Debug("error func SelectRoomList, method QueryContext by path internal/db/room/room.go",
			zap.Error(err))
		return nil, err
	}

	defer rows.Close()

	list := make([]*room.Room, 0)
	for rows.Next() {
		data := room.Room{}
		err := rows.Scan(&data.ID, &data.RoomName, &data.Title)
		if err != nil {
			r.logger.Debug("error func SelectRoomList, method Scan by path internal/db/room/room.go",
				zap.Error(err))
			continue
		}

		list = append(list, &data)
	}

	return list, nil
}
