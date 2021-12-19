package user

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	db2 "github.com/nndergunov/RTGC-Project/server/pkg/db/internal/rtgc/public"
	"github.com/nndergunov/RTGC-Project/server/pkg/db/internal/rtgc/public/model"

	// pq is needed for postgres driver.
	_ "github.com/lib/pq"
)

// usersinroom repository.

func NewUsersInRoomRepository(db db2.SQLHandler) *UsersInRoom {
	return &UsersInRoom{db}
}

type UsersInRoom struct {
	db db2.SQLHandler
}

func (u *UsersInRoom) CreateUsersInRoom(
	ctx context.Context,
	rID int, userID,
	userName string) (id int, err error) {
	roomID := int32(rID)
	usersinroom := model.Usersinroom{
		ID:       0,
		RoomID:   &roomID,
		Userid:   &userID,
		Username: &userName,
	}
	stmt := createUsersInRoomQuery(&usersinroom)
	query, args := stmt.Sql()

	row := u.db.QueryRowContext(ctx, query, args...)

	err = row.Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("scan %s result: %w", stmt.DebugSql(), err)
	}

	return
}

func (u *UsersInRoom) ReadUsersInRoom(ctx context.Context, id int) (*model.Usersinroom, error) {
	stmt := readUsersInRoomQuery(id)
	query, args := stmt.Sql()
	row := u.db.QueryRowContext(ctx, query, args...)

	users := &model.Usersinroom{
		ID:       int32(id),
		RoomID:   nil,
		Userid:   nil,
		Username: nil,
	}

	err := row.Scan(
		&users.RoomID,
		&users.Userid,
		&users.Username,
	)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return nil, db2.ErrNoData
	}

	if err != nil {
		return nil, fmt.Errorf("scan %s result: %w", stmt.DebugSql(), err)
	}

	return users, nil
}

func (u UsersInRoom) UpdateUsersInRoom(ctx context.Context, user *model.Usersinroom) error {
	stmt := updateUsersInRoomQuery(user)
	query, args := stmt.Sql()
	_, err := u.db.ExecContext(ctx, query, args...)

	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return db2.ErrNoData
	}

	if err != nil {
		return fmt.Errorf("scan %s result: %w", stmt.DebugSql(), err)
	}

	return nil
}

func (u UsersInRoom) DeleteUsersInRoom(ctx context.Context, id int) error {
	stmt := deleteUsersInRoomQuery(id)
	query, args := stmt.Sql()
	_, err := u.db.ExecContext(ctx, query, args...)

	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return db2.ErrNoData
	}

	if err != nil {
		return fmt.Errorf("scan %s result: %w", stmt.DebugSql(), err)
	}

	return nil
}

func (u *UsersInRoom) ListUsersInRoom(
	ctx context.Context,
	list *db2.ListOptions,
	crit *db2.UsersInRoomCriteria) ([]*model.Usersinroom, error) {
	stmt := listUsersInRoomQuery(list, crit)
	query, args := stmt.Sql()

	rows, err := u.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("query %s: %w", stmt.DebugSql(), err)
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Fatalf("closure error: %v", err)
		}
	}(rows)

	var res []*model.Usersinroom

	for rows.Next() {
		user := &model.Usersinroom{
			ID:       0,
			RoomID:   nil,
			Userid:   nil,
			Username: nil,
		}

		err = rows.Scan(
			&user.ID,
			&user.RoomID,
			&user.Userid,
			&user.Username,
		)
		if err != nil {
			return nil, fmt.Errorf("scan failed: %w", err)
		}

		res = append(res, user)
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}

	return res, nil
}

var _ UsersInRoomRepository = (*UsersInRoom)(nil)
