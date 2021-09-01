package checks

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_AddCheckResult_WhenCreateCheckResultWithNilError_ShouldBeHealthy(t *testing.T) {
	healthCheckResult := NewHealthCheckResult()
	checkResult := NewCheckResult("serviceName", nil)

	healthCheckResult.AddCheckResult("serviceName", checkResult)

	assert.NotNil(t, healthCheckResult)
	assert.NotNil(t, checkResult)
	assert.Equal(t, healthCheckResult.Status, Healthy)
	assert.Equal(t, len(healthCheckResult.Results), 1)

	assert.Equal(t, healthCheckResult.Results["serviceName"].Status, Healthy)
	assert.Equal(t, healthCheckResult.Results["serviceName"].Description, "serviceName is healthy")
}

func Test_AddCheckResult_WhenCreateCheckResultWithError_ShouldBeUnhealthy(t *testing.T) {
	healthCheckResult := NewHealthCheckResult()
	checkResult := NewCheckResult("serviceName", errors.New("error"))

	healthCheckResult.AddCheckResult("serviceName", checkResult)

	assert.NotNil(t, healthCheckResult)
	assert.NotNil(t, checkResult)
	assert.Equal(t, healthCheckResult.Status, Unhealthy)
	assert.Equal(t, len(healthCheckResult.Results), 1)

	assert.Equal(t, healthCheckResult.Results["serviceName"].Status, Unhealthy)
	assert.Contains(t, healthCheckResult.Results["serviceName"].Description, "ERROR:")
}
