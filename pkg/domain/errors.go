package domain

import "errors"

var (
	errID                  = errors.New("user with such ID already is in the room using valid username")
	errUname               = errors.New("such username is already taken in this room")
	errNoUser              = errors.New("no user with such userID found in a room")
	errUnsupportedUsername = errors.New("username is not supported")
	errNoRoom              = errors.New("no room with such name exists")
)
