package db

type RoomsCriteria struct {
	ID   *int
	Name *string
}

type UsersInRoomCriteria struct {
	ID       *int
	RoomID   *int
	UserID   *string
	Username *string
}
