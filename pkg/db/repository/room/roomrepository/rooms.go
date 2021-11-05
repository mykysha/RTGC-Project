package repository

import (
	"context"

	db "github.com/nndergunov/RTGC-Project/pkg/db/internal/rtgc/public"
	"github.com/nndergunov/RTGC-Project/pkg/db/internal/rtgc/public/model"
)

type RoomsRepository interface {
	CreateRooms(ctx context.Context, name string) (id int, err error)
	ReadRooms(ctx context.Context, id int) (*model.Rooms, error)
	UpdateRooms(ctx context.Context, room *model.Rooms) error
	DeleteRooms(ctx context.Context, id int) error

	ListRooms(ctx context.Context, list *db.ListOptions, criteria *db.RoomsCriteria) ([]*model.Rooms, error)
}
