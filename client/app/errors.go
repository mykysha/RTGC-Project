package app

import "errors"

// static errors.
var (
	errContain   = errors.New("unknown command: does not contain ':'")
	errSplit     = errors.New("unknown command: wrong number of arguments")
	errCom       = errors.New("unknown command")
	errUnauthUse = errors.New("unauthorised use of the command")
)
