package tests

import (
	"github.com/wesleycosta/healthcheck-go/checks/mongo"
	"github.com/wesleycostata/healthcheck-go/checks/rabbit"
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
