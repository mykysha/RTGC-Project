package jet

import (
	pg "github.com/go-jet/jet/v2/postgres"
	"github.com/nndergunov/RTGC-Project/pkg/db/internal/rtgc/public/model"
	"github.com/nndergunov/RTGC-Project/pkg/db/internal/rtgc/public/table"
	"github.com/nndergunov/RTGC-Project/pkg/db/repository"
)

// rooms query.

func createRoomsQuery(name string) pg.InsertStatement {
	return table.Rooms.INSERT(table.Rooms.Name).VALUES(name).RETURNING(table.Rooms.ID)
}

func readRoomsQuery(id int) pg.SelectStatement {
	return pg.SELECT(table.Rooms.Name).FROM(table.Rooms).WHERE(table.Rooms.ID.EQ(pg.Int(int64(id))))
}

func updateRoomsQuery(r *model.Rooms) pg.UpdateStatement {
	return table.Rooms.UPDATE(table.Rooms.Name).SET(r.Name).WHERE(table.Rooms.ID.EQ(pg.Int(int64(r.ID))))
}

func deleteRoomQuery(id int) pg.DeleteStatement {
	return table.Rooms.DELETE().WHERE(table.Rooms.ID.EQ(pg.Int(int64(id))))
}

func listRoomsQuery(listOptions *repository.ListOptions, criteria *repository.RoomsCriteria) pg.SelectStatement {
	stmt := pg.SELECT(table.Rooms.AllColumns).FROM(table.Rooms)

	conditions := pg.Bool(true)
	if criteria != nil && criteria.ID != nil {
		conditions = conditions.AND(table.Rooms.ID.EQ(pg.Int(int64(*criteria.ID))))
	}

	if criteria != nil && criteria.Name != nil {
		query := pg.RawString("'%' || query || '%'", map[string]interface{}{"query": *criteria.Name})

		conditions = conditions.AND(
			table.Rooms.Name.LIKE(query),
		)
	}

	stmt.WHERE(conditions)

	if listOptions != nil {
		for _, sort := range listOptions.Sort {
			if sort.Direction == repository.SortDirectionDESC {
				stmt.ORDER_BY(pg.RawString(sort.Property).DESC())
			} else {
				stmt.ORDER_BY(pg.RawString(sort.Property).ASC())
			}
		}
	}

	return stmt
}

// usersinrooms query.

func createUsersInRoomQuery(u *model.Usersinroom) pg.InsertStatement {
	return table.Usersinroom.INSERT(
		table.Usersinroom.RoomID,
		table.Usersinroom.Userid,
		table.Usersinroom.Username,
	).VALUES(
		*u.RoomID,
		*u.Userid,
		*u.Username,
	).RETURNING(table.Usersinroom.ID)
}

func readUsersInRoomQuery(id int) pg.SelectStatement {
	return pg.SELECT(
		table.Usersinroom.RoomID,
		table.Usersinroom.Userid,
		table.Usersinroom.Username,
	).FROM(
		table.Usersinroom,
	).WHERE(
		table.Usersinroom.ID.EQ(pg.Int(int64(id))),
	)
}

func updateUsersInRoomQuery(u *model.Usersinroom) pg.UpdateStatement {
	return table.Rooms.UPDATE(
		table.Usersinroom.RoomID,
		table.Usersinroom.Userid,
		table.Usersinroom.Username,
	).SET(
		u.RoomID,
		u.Userid,
		u.Username,
	).WHERE(table.Usersinroom.ID.EQ(pg.Int(int64(u.ID))))
}

func deleteUsersInRoomQuery(id int) pg.DeleteStatement {
	return table.Usersinroom.DELETE().WHERE(table.Usersinroom.ID.EQ(pg.Int(int64(id))))
}

func listUsersInRoomQuery(
	listOptions *repository.ListOptions,
	criteria *repository.UsersInRoomCriteria) pg.SelectStatement {
	stmt := pg.SELECT(table.Usersinroom.AllColumns).FROM(table.Usersinroom)

	conditions := pg.Bool(true)
	if criteria != nil && criteria.ID != nil {
		conditions = conditions.AND(table.Usersinroom.ID.EQ(pg.Int(int64(*criteria.ID))))
	}

	if criteria != nil && criteria.RoomID != nil {
		conditions = conditions.AND(table.Usersinroom.RoomID.EQ(pg.Int(int64(*criteria.RoomID))))
	}

	if criteria != nil && criteria.Username != nil {
		query := pg.RawString("'%' || query || '%'", map[string]interface{}{"query": *criteria.Username})

		conditions = conditions.AND(
			table.Usersinroom.Username.LIKE(query),
		)
	}

	if criteria != nil && criteria.UserID != nil {
		query := pg.RawString("'%' || query || '%'", map[string]interface{}{"query": *criteria.UserID})

		conditions = conditions.AND(
			table.Usersinroom.Userid.LIKE(query),
		)
	}

	stmt.WHERE(conditions)

	if listOptions != nil {
		for _, sort := range listOptions.Sort {
			if sort.Direction == repository.SortDirectionDESC {
				stmt.ORDER_BY(pg.RawString(sort.Property).DESC())
			} else {
				stmt.ORDER_BY(pg.RawString(sort.Property).ASC())
			}
		}
	}

	return stmt
}
