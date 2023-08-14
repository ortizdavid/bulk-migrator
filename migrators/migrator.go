package migrators

import "github.com/ortizdavid/bulk-migrator/decoders"

type Migrator struct {
}

func (migrator *Migrator) MigrateDatabase(sourceConfig string, targetConfig string) {
	var dbConfig decoders.DatabaseConfig
	source := dbConfig.DecodeDbConfig(sourceConfig)
	target := dbConfig.DecodeDbConfig(targetConfig)
	source.	PrintConfigs()
	target.PrintConfigs()
}
