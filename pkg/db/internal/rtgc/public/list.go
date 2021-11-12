package db

type SortDirection string

const (
	SortDirectionASC  SortDirection = "ASC"
	SortDirectionDESC SortDirection = "DESC"
)

type SortOrder struct {
	Property  string
	Direction SortDirection
}

type ListOptions struct {
	Sort []SortOrder
}
