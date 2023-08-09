package decoders

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type DatabaseConfig struct {
	Driver 		string `json:"driver"`
	Host 		string `json:"host"`
	Port 		int `json:"port"`
	DbName 		string `json:"db_name"`
	Tables 		[]Table `json:"tables"`
}


func (dbConfig DatabaseConfig) DecodeDbConfig(fileName string) DatabaseConfig {

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal("Error while openning file:", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&dbConfig); err != nil {
		log.Fatal("Error while decoding JSON", err)
	}

	return dbConfig
}

func (dbConfig DatabaseConfig) getNumTables() int {
	return len(dbConfig.Tables)
}

func (dbConfig DatabaseConfig) sameNumTables(source DatabaseConfig, target DatabaseConfig) bool {
	return source.getNumTables() == target.getNumTables()
}

func (dbConfig DatabaseConfig) sameNumFields(source DatabaseConfig, target DatabaseConfig) bool {
	/*for _, table := range dbConfig.Tables {
		table = ""
	}*/
	return false
}


func (dbConfig DatabaseConfig) PrintConfigs() {
	fmt.Println("Driver: ", dbConfig.Driver)
	fmt.Println("Host: ", dbConfig.Host)
	fmt.Println("Port: ", dbConfig.Port)
	fmt.Println("DB Name: ", dbConfig.DbName)
	for _, table := range dbConfig.Tables {
		table.PrintFields()
	}
}
