package room

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	// pq is needed for postgres driver.
	_ "github.com/lib/pq"
	db2 "github.com/nndergunov/RTGC-Project/server/pkg/db/internal/rtgc/public"
	"github.com/nndergunov/RTGC-Project/server/pkg/db/internal/rtgc/public/model"
)

// rooms repository.

func NewRoomsRepository(db db2.SQLHandler) *Rooms {
	return &Rooms{db}
}

type Rooms struct {
	db db2.SQLHandler
}

func (r *Rooms) CreateRooms(ctx context.Context, name string) (id int, err error) {
	stmt := createRoomsQuery(name)
	query, args := stmt.Sql()

	row := r.db.QueryRowContext(ctx, query, args...)

	err = row.Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("scan %s result: %w", stmt.DebugSql(), err)
	}

	return
}

func (r *Rooms) ReadRooms(ctx context.Context, id int) (*model.Rooms, error) {
	stmt := readRoomsQuery(id)
	query, args := stmt.Sql()
	row := r.db.QueryRowContext(ctx, query, args...)

	rooms := &model.Rooms{
		ID:   int32(id),
		Name: nil,
	}

	err := row.Scan(&rooms.Name)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return nil, db2.ErrNoData
	}

	if err != nil {
		return nil, fmt.Errorf("scan %s result: %w", stmt.DebugSql(), err)
	}

	return rooms, nil
}

func (r Rooms) UpdateRooms(ctx context.Context, room *model.Rooms) error {
	stmt := updateRoomsQuery(room)
	query, args := stmt.Sql()
	_, err := r.db.ExecContext(ctx, query, args...)

	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return db2.ErrNoData
	}

	if err != nil {
		return fmt.Errorf("scan %s result: %w", stmt.DebugSql(), err)
	}

	return nil
}

func (r Rooms) DeleteRooms(ctx context.Context, id int) error {
	stmt := deleteRoomQuery(id)
	query, args := stmt.Sql()
	_, err := r.db.ExecContext(ctx, query, args...)

	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return db2.ErrNoData
	}

	if err != nil {
		return fmt.Errorf("scan %s result: %w", stmt.DebugSql(), err)
	}

	return nil
}

func (r *Rooms) ListRooms(
	ctx context.Context,
	list *db2.ListOptions,
	crit *db2.RoomsCriteria) ([]*model.Rooms, error) {
	stmt := listRoomsQuery(list, crit)
	query, args := stmt.Sql()

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("query %s: %w", stmt.DebugSql(), err)
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Fatalf("closure error: %v", err)
		}
	}(rows)

	var res []*model.Rooms

	for rows.Next() {
		room := &model.Rooms{
			ID:   0,
			Name: nil,
		}

		err = rows.Scan(
			&room.ID,
			&room.Name,
		)
		if err != nil {
			return nil, fmt.Errorf("scan failed: %w", err)
		}

		res = append(res, room)
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}

	return res, nil
}

var _ RoomsRepository = (*Rooms)(nil)
