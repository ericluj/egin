package egin

type ErrorType uint64

type Error struct {
	Err  error
	Type ErrorType
	Meta any
}

type errorMsgs []*Error
