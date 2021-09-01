package healthcheck

import (
	"testing"

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

func Test_Execute_WhenRabbitWithInvalidConfiguration_ShouldReturnHealthy(t *testing.T) {
	healthCheck := New()
	healthCheck.AddRabbit(&rabbit.Config{})
	healthCheckResult := healthCheck.Execute()

	assert.NotNil(t, healthCheckResult)
	assert.Equal(t, healthCheckResult.Status, check.Unhealthy)
	assert.Equal(t, len(healthCheckResult.Results), 1)
	assert.Contains(t, healthCheckResult.Results["rabbit"].Description, "ERROR:")
}
