package healthcheck

import (
	"testing"

	tests "github.com/mundipagg/healthcheck-go/tests"

	check "github.com/mundipagg/healthcheck-go/checks"
	rabbit "github.com/mundipagg/healthcheck-go/checks/rabbit"
	"github.com/stretchr/testify/assert"
)

func Test_Execute_WhenNoCheckConfigured_ShouldReturnHealthy(t *testing.T) {
	healthCheck := New()

	healthCheckResult := healthCheck.Execute()

	assert.NotNil(t, healthCheckResult)
	assert.Equal(t, healthCheckResult.Status, check.Healthy)
	assert.Equal(t, len(healthCheckResult.Results), 0)
}

func Test_Execute_WhenRabbitWithInvalidConfiguration_ShouldReturnUnhealthy(t *testing.T) {
	healthCheck := New()
	config := &rabbit.Config{}
	healthCheck.AddService(config)
	healthCheckResult := healthCheck.Execute()

	assert.NotNil(t, healthCheckResult)
	assert.Equal(t, healthCheckResult.Status, check.Unhealthy)
	assert.Equal(t, len(healthCheckResult.Results), 1)
	assert.Contains(t, healthCheckResult.Results["rabbit"].Description, "ERROR:")
}

func Test_Execute_WhenRabbitAndMongoIsRunningWithValidConfiguration_ShouldReturnHealthy(t *testing.T) {
	healthCheck := New()

	configRabbit := tests.NewStubRabbitConfig()
	configMongo := tests.NewStubMongoConfig()

	healthCheck.AddService(configRabbit)
	healthCheck.AddService(configMongo)

	healthCheckResult := healthCheck.Execute()

	assert.NotNil(t, healthCheckResult)
	assert.Equal(t, healthCheckResult.Status, check.Healthy)
	assert.Equal(t, len(healthCheckResult.Results), 2)
	assert.Contains(t, healthCheckResult.Results["rabbit"].Status, check.Healthy)
	assert.Contains(t, healthCheckResult.Results["rabbit"].Description, "rabbit is healthy")

	assert.Contains(t, healthCheckResult.Results["mongo"].Status, check.Healthy)
	assert.Contains(t, healthCheckResult.Results["mongo"].Description, "mongo is healthy")
}
