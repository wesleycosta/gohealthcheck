package mongo

import (
	"testing"

	check "github.com/mundipagg/healthcheck-go/checks"
	"github.com/stretchr/testify/assert"
)

func Test_GetName_WhenExecuted_ShouldReturnMongo(t *testing.T) {
	config := newStubMongoConfig()
	healthcheck := config.CreateCheck()

	assert.Equal(t, healthcheck.GetName(), "mongo")
}

func Test_Execute_WhenConfigurationIsValidAndServiceIsRunning_ShouldReturnHealthy(t *testing.T) {
	config := newStubMongoConfig()

	healthcheck := config.CreateCheck()
	result := healthcheck.Execute()

	assert.NotNil(t, result)
	assert.Equal(t, result.Status, check.Healthy)
	assert.Equal(t, result.Description, "mongo is healthy")
}

func Test_Execute_WhenConfigurationIsInvalid_ShouldReturnUnhealthy(t *testing.T) {
	invalidConfig := newStubMongoConfig()
	invalidConfig.withUrl("")

	healthCheck := invalidConfig.CreateCheck()
	result := healthCheck.Execute()

	assert.NotNil(t, result)
	assert.Equal(t, result.Status, check.Unhealthy)
	assert.Contains(t, result.Description, "ERROR:")
}
