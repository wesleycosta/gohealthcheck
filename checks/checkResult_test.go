package checks

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewCheckResult_WhenCreateCheckResultWithNilError_ShouldBeHealthy(t *testing.T) {
	result := NewCheckResult("serviceName", nil)

	assert.NotNil(t, result)
	assert.Equal(t, result.Status, Healthy)
	assert.Equal(t, result.Description, "serviceName is healthy")
}

func Test_NewCheckResult_WhenCreateCheckResultWithError_ShouldBeUnhealthy(t *testing.T) {
	result := NewCheckResult("serviceName", errors.New("error"))

	assert.NotNil(t, result)
	assert.Equal(t, result.Status, Unhealthy)
	assert.Contains(t, result.Description, "ERROR:")
}
