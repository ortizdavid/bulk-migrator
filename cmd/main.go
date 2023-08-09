package main

import (
	"fmt"
	"os"

	"github.com/ortizdavid/bulk-migrator/decoders"
)

func main() {
	fmt.Println(os.Args)
	
	var dbConfig decoders.DatabaseConfig

	source := dbConfig.DecodeDbConfig("../samples/source.json")
	target := dbConfig.DecodeDbConfig("../samples/target.json")
	
	source.PrintConfigs()
	target.PrintConfigs()
}