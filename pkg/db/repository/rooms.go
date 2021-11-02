package repository

import (
	"context"

	"github.com/nndergunov/RTGC-Project/pkg/db/internal/rtgc/public/model"
)

type RoomsCriteria struct {
	ID   *int
	Name *string
}

type RoomsRepository interface {
	CreateRooms(ctx context.Context, name string) (id int, err error)
	ReadRooms(ctx context.Context, id int) (*model.Rooms, error)
	UpdateRooms(ctx context.Context, room *model.Rooms) error
	DeleteRooms(ctx context.Context, id int) error

	ListRooms(ctx context.Context, list *ListOptions, criteria *RoomsCriteria) ([]*model.Rooms, error)
}
