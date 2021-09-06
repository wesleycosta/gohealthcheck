package sqlServer

import (
	"testing"

	check "github.com/mundipagg/healthcheck-go/checks"
	"github.com/stretchr/testify/assert"
)

func Test_GetName_WhenExecuted_ShouldReturnSqlServer(t *testing.T) {
	config := newStubSqlServerConfig()
	healthcheck := config.CreateCheck()

	assert.Equal(t, healthcheck.GetName(), "sqlServer")
}

func Test_Execute_WhenConfigurationIsValidAndServiceIsRunning_ShouldReturnHealthy(t *testing.T) {
	config := newStubSqlServerConfig()
	healthcheck := config.CreateCheck()
	result := healthcheck.Execute()

	assert.NotNil(t, result)
	assert.Equal(t, result.Status, check.Healthy)
	assert.Equal(t, result.Description, "sqlServer is healthy")
}

func Test_Execute_WhenConfigurationIsInvalid_ShouldReturnUnhealthy(t *testing.T) {
	invalidConfig := newStubSqlServerConfig()
	invalidConfig.withQuery("invalidQuery")

	healthcheck := invalidConfig.CreateCheck()
	result := healthcheck.Execute()

	assert.NotNil(t, result)
	assert.Equal(t, result.Status, check.Unhealthy)
	assert.Contains(t, result.Description, "ERROR:")
}
