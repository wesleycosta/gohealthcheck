package mongo

import (
	"context"
	"crypto/tls"
	"fmt"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

var (
	conn              *mongo.Client // is concurrent safe: https://github.com/mongodb/mongo-go-driver/blob/master/mongo/client.go#L46
	ConnectionTimeout = 3 * time.Second
	mu                sync.RWMutex
)

const (
	NotFoundDoc = "mongo: no documents in result"
	InvalidPK   = "invalid pk"
	emptyConn   = "Connection is empty"
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
	Url        string
	User       string
	Password   string
	Database   string
	AuthSource string
	Timeout    int
	ForceTLS   bool
}

type healthCheck struct {
	Config *Config
}

func (service *healthCheck) GetName() string {
	return "mongo"
}

func (healthCheck *healthCheck) Execute() error {
	_, err := healthCheck.createMongo()
	if err != nil {
		return err
	}

	return ping()
}

func (healthCheck *healthCheck) createMongo() (*mongo.Client, error) {
	mu.Lock()
	defer mu.Unlock()

	if conn != nil {
		return conn, nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), ConnectionTimeout)
	defer cancel()

	var err error
	conn, err = mongo.Connect(ctx, getClientOptions(healthCheck.Config))
	if err != nil {
		return conn, err
	}

	return conn, nil
}

func ping() error {
	if conn == nil {
		return fmt.Errorf(emptyConn)
	}

	ctx, cancel := context.WithTimeout(context.Background(), ConnectionTimeout)
	defer cancel()

	err := conn.Ping(ctx, readpref.Primary())
	if err != nil {
		return err
	}

	return nil
}

func getClientOptions(config *Config) *options.ClientOptions {
	mongoURL := config.Url
	co := options.Client()
	co.SetRetryWrites(true)
	co.SetWriteConcern(writeconcern.New(writeconcern.WMajority()))

	co.SetConnectTimeout(time.Duration(config.Timeout) * time.Second)
	co.SetMaxConnIdleTime(time.Duration(config.Timeout) * time.Second)
	co.SetMaxPoolSize(512)

	if config.ForceTLS {
		co.SetTLSConfig(&tls.Config{})
	}

	return co.ApplyURI(mongoURL).SetAuth(mongoCredential(config))
}

func mongoCredential(config *Config) options.Credential {
	user := config.User
	password := config.Password
	var database string
	if config.AuthSource != "" {
		database = config.AuthSource
	} else {
		database = config.Database
	}

	credential := options.Credential{
		Username:   user,
		Password:   password,
		AuthSource: database,
	}

	if config.ForceTLS {
		credential.AuthMechanism = "SCRAM-SHA-1"
	}

	return credential
}
