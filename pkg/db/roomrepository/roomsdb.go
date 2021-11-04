package roomrepository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	// pq is needed for postgres driver.
	_ "github.com/lib/pq"
	db "github.com/nndergunov/RTGC-Project/pkg/db/internal/rtgc/public"
	"github.com/nndergunov/RTGC-Project/pkg/db/internal/rtgc/public/model"
	repo "github.com/nndergunov/RTGC-Project/pkg/db/roomrepository/repository"
)

// rooms repository.

func NewRoomsRepository(db db.SQLHandler) *RoomsRepository {
	return &RoomsRepository{db}
}

type RoomsRepository struct {
	db db.SQLHandler
}

func (r *RoomsRepository) CreateRooms(ctx context.Context, name string) (id int, err error) {
	stmt := createRoomsQuery(name)
	query, args := stmt.Sql()

	row := r.db.QueryRowContext(ctx, query, args...)

	err = row.Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("scan %s result: %w", stmt.DebugSql(), err)
	}

	return
}

func (r *RoomsRepository) ReadRooms(ctx context.Context, id int) (*model.Rooms, error) {
	stmt := readRoomsQuery(id)
	query, args := stmt.Sql()
	row := r.db.QueryRowContext(ctx, query, args...)

	rooms := &model.Rooms{
		ID:   int32(id),
		Name: nil,
	}

	err := row.Scan(&rooms.Name)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return nil, db.ErrNoData
	}

	if err != nil {
		return nil, fmt.Errorf("scan %s result: %w", stmt.DebugSql(), err)
	}

	return rooms, nil
}

func (r RoomsRepository) UpdateRooms(ctx context.Context, room *model.Rooms) error {
	stmt := updateRoomsQuery(room)
	query, args := stmt.Sql()
	_, err := r.db.ExecContext(ctx, query, args...)

	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return db.ErrNoData
	}

	if err != nil {
		return fmt.Errorf("scan %s result: %w", stmt.DebugSql(), err)
	}

	return nil
}

func (r RoomsRepository) DeleteRooms(ctx context.Context, id int) error {
	stmt := deleteRoomQuery(id)
	query, args := stmt.Sql()
	_, err := r.db.ExecContext(ctx, query, args...)

	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return db.ErrNoData
	}

	if err != nil {
		return fmt.Errorf("scan %s result: %w", stmt.DebugSql(), err)
	}

	return nil
}

func (r *RoomsRepository) ListRooms(
	ctx context.Context,
	list *db.ListOptions,
	crit *db.RoomsCriteria) ([]*model.Rooms, error) {
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

var _ repo.RoomsRepository = (*RoomsRepository)(nil)
