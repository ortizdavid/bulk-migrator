package helpers

import "errors"

var (
	ErrorNumTables = errors.New("origin and destiny must be same number of tables")
	ErrorNumFields = errors.New("origin and destiny must be same number of fields")
)