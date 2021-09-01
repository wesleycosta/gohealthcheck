package tests

import (
	"github.com/mundipagg/healthcheck-go/checks/mongo"
	"github.com/mundipagg/healthcheck-go/checks/rabbit"
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
