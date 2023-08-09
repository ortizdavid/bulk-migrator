package loaders

import "github.com/ortizdavid/bulk-migrator/decoders"

type Loader struct {
}

func (loader *Loader) LoadDatabase(sourceConfig string, targetConfig string) {
	var dbConfig decoders.DatabaseConfig
	source := dbConfig.DecodeDbConfig(sourceConfig)
	target := dbConfig.DecodeDbConfig(targetConfig)
	source.	PrintConfigs()
	target.PrintConfigs()
}
