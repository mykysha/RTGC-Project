package roomrepository

import (
	pg "github.com/go-jet/jet/v2/postgres"
	db "github.com/nndergunov/RTGC-Project/pkg/db/internal/rtgc/public"
	"github.com/nndergunov/RTGC-Project/pkg/db/internal/rtgc/public/model"
	"github.com/nndergunov/RTGC-Project/pkg/db/internal/rtgc/public/table"
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

func listRoomsQuery(listOptions *db.ListOptions, criteria *db.RoomsCriteria) pg.SelectStatement {
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
			if sort.Direction == db.SortDirectionDESC {
				stmt.ORDER_BY(pg.RawString(sort.Property).DESC())
			} else {
				stmt.ORDER_BY(pg.RawString(sort.Property).ASC())
			}
		}
	}

	return stmt
}
