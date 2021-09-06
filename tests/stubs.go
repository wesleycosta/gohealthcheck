package tests

import (
	"github.com/wesleycosta/healthcheck-go/checks/mongo"
	"github.com/wesleycosta/healthcheck-go/checks/rabbit"
	"github.com/wesleycosta/healthcheck-go/checks/sqlServer"
)

func NewStubRabbitConfig() *rabbit.Config {
	return &rabbit.Config{
		ConnectionString: "amqp://guest:guest@localhost:5672/",
	}
}

func NewStubMongoConfig() *mongo.Config {
	return &mongo.Config{
		Url:        "mongodb://localhost:27017",
		User:       "test",
		Password:   "test",
		AuthSource: "admin",
		Timeout:    3,
		ForceTLS:   false,
	}
}

func NewStubSqlServerConfig() *sqlServer.Config {
	return &sqlServer.Config{
		ConnectionString: "server=localhost;port=1434;user id=sa;password=sa;database=master;connection timeout=130",
		Query:            "SELECT TOP 1 TABLE_NAME from INFORMATION_SCHEMA.TABLES",
	}
}
