package healthcheck

import (
	"testing"

	tests "github.com/wesleycosta/healthcheck-go/tests"

	"github.com/stretchr/testify/assert"
	check "github.com/wesleycosta/healthcheck-go/checks"
	rabbit "github.com/wesleycosta/healthcheck-go/checks/rabbit"
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

func Test_Execute_WhenDependenciesIsRunningWithValidConfiguration_ShouldReturnHealthy(t *testing.T) {
	healthCheck := New()

	configRabbit := tests.NewStubRabbitConfig()
	configMongo := tests.NewStubMongoConfig()
	configSqlServer := tests.NewStubSqlServerConfig()

	healthCheck.AddService(configRabbit)
	healthCheck.AddService(configMongo)
	healthCheck.AddService(configSqlServer)

	healthCheckResult := healthCheck.Execute()

	assert.NotNil(t, healthCheckResult)
	assert.Equal(t, healthCheckResult.Status, check.Healthy)
	assert.Equal(t, len(healthCheckResult.Results), 3)
	assert.Contains(t, healthCheckResult.Results["rabbit"].Status, check.Healthy)
	assert.Contains(t, healthCheckResult.Results["rabbit"].Description, "rabbit is healthy")

	assert.Contains(t, healthCheckResult.Results["mongo"].Status, check.Healthy)
	assert.Contains(t, healthCheckResult.Results["mongo"].Description, "mongo is healthy")

	assert.Contains(t, healthCheckResult.Results["sqlServer"].Status, check.Healthy)
	assert.Contains(t, healthCheckResult.Results["sqlServer"].Description, "sqlServer is healthy")
}
