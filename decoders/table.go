package decoders

import "fmt"

type Table struct {
	TableName string   `json:"table_name"`
	Fields    []string `json:"fields"`
}

func (table Table) PrintFields() {
	fmt.Println("Table: ", table.TableName)
	fmt.Println("Fields: ")
	for _, field := range table.Fields {
		fmt.Println("\t", field)
	}
}

func (table Table) getNumFields() int {
	return len(table.Fields)
}

