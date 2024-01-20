package room

import (
	"context"
	"database/sql"
	roomEntity "github.com/EvgeniyBudaev/golang-clean-architecture-evrone/internal/entity/room"
	"github.com/EvgeniyBudaev/golang-clean-architecture-evrone/internal/logger"
	"go.uber.org/zap"
)

type DBRoom interface {
	CreateRoom(ctx *context.Context, room *roomEntity.Room) (*roomEntity.Room, error)
	SelectRoomList(ctx *context.Context, room *roomEntity.Room) ([]*roomEntity.Room, error)
}

type PGRoomDB struct {
	db *sql.DB
}

func NewPGRoomDB(db *sql.DB) *PGRoomDB {
	return &PGRoomDB{db: db}
}

func (pg *PGRoomDB) CreateRoom(ctx context.Context, r *roomEntity.Room) (*roomEntity.Room, error) {
	query := "INSERT INTO rooms (room_name, title) VALUES ($1, $2) RETURNING id"
	err := pg.db.QueryRowContext(ctx, query, r.RoomName, r.Title).Scan(&r.ID)
	if err != nil {
		logger.Log.Debug("error func CreateRoom, method QueryRowContext by path internal/db/room/room.go",
			zap.Error(err))
		return nil, err
	}
	return r, nil
}

func (pg *PGRoomDB) SelectRoomList(ctx context.Context) ([]*roomEntity.Room, error) {
	query := "SELECT id, room_name, title FROM rooms"
	rows, err := pg.db.QueryContext(ctx, query)
	if err != nil {
		logger.Log.Debug("error func SelectRoomList, method QueryContext by path internal/db/room/room.go",
			zap.Error(err))
		return nil, err
	}
	defer rows.Close()
	list := make([]*roomEntity.Room, 0)
	for rows.Next() {
		data := roomEntity.Room{}
		err := rows.Scan(&data.ID, &data.RoomName, &data.Title)
		if err != nil {
			logger.Log.Debug("error func SelectRoomList, method Scan by path internal/db/room/room.go",
				zap.Error(err))
			continue
		}
		list = append(list, &data)
	}
	return list, nil
}
