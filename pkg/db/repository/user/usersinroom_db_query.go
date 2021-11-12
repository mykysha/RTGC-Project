package user

import (
	pg "github.com/go-jet/jet/v2/postgres"
	db "github.com/nndergunov/RTGC-Project/pkg/db/internal/rtgc/public"
	"github.com/nndergunov/RTGC-Project/pkg/db/internal/rtgc/public/model"
	"github.com/nndergunov/RTGC-Project/pkg/db/internal/rtgc/public/table"
)

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
	listOptions *db.ListOptions,
	criteria *db.UsersInRoomCriteria) pg.SelectStatement {
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
			if sort.Direction == db.SortDirectionDESC {
				stmt.ORDER_BY(pg.RawString(sort.Property).DESC())
			} else {
				stmt.ORDER_BY(pg.RawString(sort.Property).ASC())
			}
		}
	}

	return stmt
}
