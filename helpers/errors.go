package helpers

import "errors"

var (
	ErrNumTables = errors.New("origin and destiny must be same number of tables")
	ErrNumFields = errors.New("origin and destiny must be same number of fields")
)