package sqlServer

func newStubSqlServerConfig() *Config {
	return &Config{
		ConnectionString: "teste",
		Query:            "SELECT TOP 1 TABLE_NAME from INFORMATION_SCHEMA.TABLES",
	}
}

func (config *Config) withQuery(query string) *Config {
	config.Query = query

	return config
}
