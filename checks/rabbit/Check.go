package rabbit

import (
	"errors"
	"time"

	"github.com/streadway/amqp"
)

type HealthCheck interface {
	Execute() error
	GetName() string
}

func New(config *Config) HealthCheck {
	return &healthCheck{
		Config: config,
	}
}

type Config struct {
	ConnectionString string
}

type healthCheck struct {
	Config *Config
}

func (service *healthCheck) GetName() string {
	return "rabbit"
}

func (service *healthCheck) Execute() error {
	conn, err := openConnection(service.Config)
	if err != nil {
		return err
	}

	channel, err := openChannel(conn)
	if err != nil {
		return err
	}

	return closeChannel(channel)
}

func openChannel(conn *amqp.Connection) (*amqp.Channel, error) {
	var channel *amqp.Channel
	var err error

	if conn == nil {
		err = errors.New("failed to open a channel. The connection is closed")
		return nil, err
	}

	if channel, err = conn.Channel(); err != nil {
		return nil, err
	}

	return channel, err
}

func closeChannel(channel *amqp.Channel) error {
	if channel != nil {
		return channel.Close()
	}

	return nil
}

func openConnection(config *Config) (*amqp.Connection, error) {
	var hb int
	var err error

	conn, err := amqp.DialConfig(config.ConnectionString, amqp.Config{
		Heartbeat: time.Duration(hb) * time.Second,
	})

	return conn, err
}
