package repository

import (
	"context"

	"github.com/nndergunov/RTGC-Project/pkg/db/internal/rtgc/public/model"
)

type UsersInRoomCriteria struct {
	ID       *int
	RoomID   *int
	UserID   *string
	Username *string
}

type UsersInRoomRepository interface {
	CreateUsersInRoom(ctx context.Context, roomID int, userName, userID string) (id int, err error)
	ReadUsersInRoom(ctx context.Context, id int) (*model.Usersinroom, error)
	UpdateUsersInRoom(ctx context.Context, room *model.Usersinroom) error
	DeleteUsersInRoom(ctx context.Context, id int) error

	ListUsersInRoom(ctx context.Context, list *ListOptions, criteria *UsersInRoomCriteria) ([]*model.Usersinroom, error)
}
