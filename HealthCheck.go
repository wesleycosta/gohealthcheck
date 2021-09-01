package healthcheck

import (
	"github.com/mundipagg/healthcheck-go/checks"
	mongo "github.com/mundipagg/healthcheck-go/checks/mongo"
	rabbit "github.com/mundipagg/healthcheck-go/checks/rabbit"
)

type HealthCheck interface {
	Execute() checks.HealthCheckResult
	AddRabbit(config *rabbit.Config)
	AddMongo(config *mongo.Config)
}

type healthCheck struct {
	checks map[string]checks.Check
}

func New() HealthCheck {
	return &healthCheck{
		checks: make(map[string]checks.Check),
	}
}

func (healthCheck *healthCheck) Execute() checks.HealthCheckResult {
	healthCheckResult := checks.NewHealthCheckResult()

	for key := range healthCheck.checks {
		healthCheckResult.AddCheckResult(key, healthCheck.checks[key].Execute())
	}

	return healthCheckResult
}

func (healthCheck *healthCheck) AddRabbit(config *rabbit.Config) {
	rabbitCheck := rabbit.New(config)
	healthCheck.checks[rabbitCheck.GetName()] = rabbitCheck
}

func (healthCheck *healthCheck) AddMongo(config *mongo.Config) {
	mongoCheck := mongo.New(config)
	healthCheck.checks[mongoCheck.GetName()] = mongoCheck
}
