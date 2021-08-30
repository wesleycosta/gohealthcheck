package health

import (
	"fmt"

	mongo "github.com/mundipagg/boleto-api/health/checks/mongo"
	rabbit "github.com/mundipagg/boleto-api/health/checks/rabbit"
)

const (
	Unhealthy string = "Unhealthy"
	Healthy   string = "Healthy"
)

type CheckResult struct {
	Status      string `json:"status"`
	Description string `json:"description"`
}

type HealthCheckResult struct {
	Status  string                 `json:"status"`
	Results map[string]CheckResult `json:"results"`
}

func newCheckResult() CheckResult {
	return CheckResult{
		Status: Healthy,
	}
}

func newHealthCheckResult() HealthCheckResult {
	return HealthCheckResult{
		Results: make(map[string]CheckResult),
	}
}

func (checkResult *CheckResult) setUnhealthy(err error) {
	checkResult.Status = Unhealthy
	checkResult.Description = fmt.Sprintf("ERROR: %s", err)
}

func (healthCheckResult *HealthCheckResult) addCheckResult(key string, checkResult CheckResult) {
	if checkResult.Status == Healthy {
		checkResult.Description = fmt.Sprintf("%s is healthy", key)
	}

	healthCheckResult.Results[key] = checkResult
}

func (healthCheckResult *HealthCheckResult) setStatus() {
	for key := range healthCheckResult.Results {
		if healthCheckResult.Results[key].Status == Unhealthy {
			healthCheckResult.Status = Unhealthy
			return
		}
	}

	healthCheckResult.Status = Healthy
}

type HealthCheck interface {
	Execute() HealthCheckResult
	AddRabbit(config *rabbit.Config)
	AddMongo(config *mongo.Config)
}

type healthCheck struct {
	mongo  mongo.HealthCheck
	rabbit rabbit.HealthCheck
}

func New() HealthCheck {
	return &healthCheck{}
}

func (healthCheck *healthCheck) Execute() HealthCheckResult {
	healthCheckResult := newHealthCheckResult()

	if healthCheck.mongoExists() {
		mongoResult := healthCheck.executeMongoCheck()
		healthCheckResult.addCheckResult(healthCheck.mongo.GetName(), mongoResult)
	}

	if healthCheck.rabbitExists() {
		rabbitResult := healthCheck.executeRabbitCheck()
		healthCheckResult.addCheckResult(healthCheck.rabbit.GetName(), rabbitResult)
	}

	healthCheckResult.setStatus()
	return healthCheckResult
}

func (healthCheck *healthCheck) AddRabbit(config *rabbit.Config) {
	healthCheck.rabbit = rabbit.New(config)
}

func (healthCheck *healthCheck) AddMongo(config *mongo.Config) {
	healthCheck.mongo = mongo.New(config)
}

func (healthCheck *healthCheck) mongoExists() bool {
	return healthCheck.mongo != nil
}

func (healthCheck *healthCheck) rabbitExists() bool {
	return healthCheck.rabbit != nil
}

func (healthCheck *healthCheck) executeMongoCheck() CheckResult {
	rabbitResult := newCheckResult()
	if err := healthCheck.mongo.Execute(); err != nil {
		rabbitResult.setUnhealthy(err)
	}

	return rabbitResult
}

func (healthCheck *healthCheck) executeRabbitCheck() CheckResult {
	rabbitResult := newCheckResult()
	if err := healthCheck.rabbit.Execute(); err != nil {
		rabbitResult.setUnhealthy(err)
	}

	return rabbitResult
}
