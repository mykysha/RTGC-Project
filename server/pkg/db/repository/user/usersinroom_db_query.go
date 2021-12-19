package user

import (
	pg "github.com/go-jet/jet/v2/postgres"
	db2 "github.com/nndergunov/RTGC-Project/server/pkg/db/internal/rtgc/public"
	"github.com/nndergunov/RTGC-Project/server/pkg/db/internal/rtgc/public/model"
	table2 "github.com/nndergunov/RTGC-Project/server/pkg/db/internal/rtgc/public/table"
)

// usersinrooms query.

func createUsersInRoomQuery(u *model.Usersinroom) pg.InsertStatement {
	return table2.Usersinroom.INSERT(
		table2.Usersinroom.RoomID,
		table2.Usersinroom.Userid,
		table2.Usersinroom.Username,
	).VALUES(
		*u.RoomID,
		*u.Userid,
		*u.Username,
	).RETURNING(table2.Usersinroom.ID)
}

func readUsersInRoomQuery(id int) pg.SelectStatement {
	return pg.SELECT(
		table2.Usersinroom.RoomID,
		table2.Usersinroom.Userid,
		table2.Usersinroom.Username,
	).FROM(
		table2.Usersinroom,
	).WHERE(
		table2.Usersinroom.ID.EQ(pg.Int(int64(id))),
	)
}

func updateUsersInRoomQuery(u *model.Usersinroom) pg.UpdateStatement {
	return table2.Rooms.UPDATE(
		table2.Usersinroom.RoomID,
		table2.Usersinroom.Userid,
		table2.Usersinroom.Username,
	).SET(
		u.RoomID,
		u.Userid,
		u.Username,
	).WHERE(table2.Usersinroom.ID.EQ(pg.Int(int64(u.ID))))
}

func deleteUsersInRoomQuery(id int) pg.DeleteStatement {
	return table2.Usersinroom.DELETE().WHERE(table2.Usersinroom.ID.EQ(pg.Int(int64(id))))
}

func listUsersInRoomQuery(
	listOptions *db2.ListOptions,
	criteria *db2.UsersInRoomCriteria) pg.SelectStatement {
	stmt := pg.SELECT(table2.Usersinroom.AllColumns).FROM(table2.Usersinroom)

	conditions := pg.Bool(true)
	if criteria != nil && criteria.ID != nil {
		conditions = conditions.AND(table2.Usersinroom.ID.EQ(pg.Int(int64(*criteria.ID))))
	}

	if criteria != nil && criteria.RoomID != nil {
		conditions = conditions.AND(table2.Usersinroom.RoomID.EQ(pg.Int(int64(*criteria.RoomID))))
	}

	if criteria != nil && criteria.Username != nil {
		query := pg.RawString("'%' || query || '%'", map[string]interface{}{"query": *criteria.Username})

		conditions = conditions.AND(
			table2.Usersinroom.Username.LIKE(query),
		)
	}

	if criteria != nil && criteria.UserID != nil {
		query := pg.RawString("'%' || query || '%'", map[string]interface{}{"query": *criteria.UserID})

		conditions = conditions.AND(
			table2.Usersinroom.Userid.LIKE(query),
		)
	}

	stmt.WHERE(conditions)

	if listOptions != nil {
		for _, sort := range listOptions.Sort {
			if sort.Direction == db2.SortDirectionDESC {
				stmt.ORDER_BY(pg.RawString(sort.Property).DESC())
			} else {
				stmt.ORDER_BY(pg.RawString(sort.Property).ASC())
			}
		}
	}

	return stmt
}
