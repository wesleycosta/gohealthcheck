package rabbit

import (
	"testing"

	"github.com/stretchr/testify/assert"
	check "github.com/wesleycosta/healthcheck-go/checks"
)

func Test_GetName_WhenExecuted_ShouldReturnRabbit(t *testing.T) {
	config := newStubRabbitConfig()
	healthcheck := config.CreateCheck()

	assert.Equal(t, healthcheck.GetName(), "rabbit")
}

func Test_Execute_WhenConfigurationIsValidAndServiceIsRunning_ShouldReturnHealthy(t *testing.T) {
	config := newStubRabbitConfig()
	healthcheck := config.CreateCheck()
	result := healthcheck.Execute()

	assert.NotNil(t, result)
	assert.Equal(t, result.Status, check.Healthy)
	assert.Equal(t, result.Description, "rabbit is healthy")
}

func Test_Execute_WhenConfigurationIsInvalid_ShouldReturnUnhealthy(t *testing.T) {
	invalidConfig := newStubRabbitConfig()
	invalidConfig.withConnectionString("amqp://guest:guest@localhost:11122/")

	healthcheck := invalidConfig.CreateCheck()
	result := healthcheck.Execute()

	assert.NotNil(t, result)
	assert.Equal(t, result.Status, check.Unhealthy)
	assert.Contains(t, result.Description, "ERROR:")
}
