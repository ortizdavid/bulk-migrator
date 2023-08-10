package main

import (
	"os"
	"github.com/ortizdavid/bulk-migrator/helpers"
	"github.com/ortizdavid/bulk-migrator/migrators"
)

func main() {
	
	cliArgs := os.Args
	numArgs := len(cliArgs)

	if numArgs <= 1 {
		helpers.PrintUsage()

	} else {
		
		if numArgs == 2 &&  cliArgs[1] == "-help" {
			helpers.PrintHelp()

		} else if numArgs == 2 &&  cliArgs[1] == "-examples" {
			helpers.PrintExamples()

		} else if numArgs == 5 {
			var migrator migrators.Migrator
			sourceConfig := cliArgs[2]
			targetConfig := cliArgs[4]
			migrator.MigrateDatabase(sourceConfig, targetConfig)
			
		} else {
			helpers.PrintHelp()
		}
	}
}