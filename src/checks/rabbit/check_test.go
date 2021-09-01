package rabbit

import (
	"testing"

	check "github.com/mundipagg/healthcheck-go/checks"
	"github.com/stretchr/testify/assert"
)

func Test_GetName_WhenExecuted_ShouldReturnRabbit(t *testing.T) {
	healthcheck := New(nil)

	assert.Equal(t, healthcheck.GetName(), "rabbit")
}

func Test_Execute_WhenConfigurationIsValidAndServiceIsRunning_ShouldReturnHealthy(t *testing.T) {
	config := newStubRabbitConfig()
	healthcheck := New(config)
	result := healthcheck.Execute()

	assert.NotNil(t, result)
	assert.Equal(t, result.Status, check.Healthy)
	assert.Equal(t, result.Description, "rabbit is healthy")
}

func Test_Execute_WhenConfigurationIsInvalid_ShouldReturnUnhealthy(t *testing.T) {
	invalidConfig := newStubRabbitConfig()
	invalidConfig.withConnectionString("amqp://guest:guest@localhost:11122/")

	healthcheck := New(invalidConfig)
	result := healthcheck.Execute()

	assert.NotNil(t, result)
	assert.Equal(t, result.Status, check.Unhealthy)
	assert.Contains(t, result.Description, "ERROR:")
}
