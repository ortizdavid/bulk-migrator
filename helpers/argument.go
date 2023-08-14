package helpers

import "fmt"

type Argument struct {
	Name        string
	Description string
	ArgType     string
}

func (argument *Argument) GetFlags() []Argument {
	return []Argument{
		{ Name: "-from", Description: "Source Configuration" },
		{ Name: "-to", Description: "Target configuration" },
		{ Name: "-examples", Description: "Examples of source/target" },
		{ Name: "-help", Description: "Help for Usage" },
		{ Name: "-drivers", Description: "List all Database Drivers" },
	}
}

func (argument *Argument) Contains(args []Argument, name string) bool {
	for _, arg := range args {
		if name == arg.Name {
			return true
		}
	}
	return false
}

func (argument *Argument) PrintArguments(args []Argument) {
	for _, arg := range args {
		fmt.Printf("\t%s ----- %s\n", arg.Name, arg.Description)
	}
}