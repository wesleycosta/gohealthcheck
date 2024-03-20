# Health Check Go

Health Check library in Golang.

## Project Structure

- **checks:** Folder containing checks for external dependencies (SQL Server, RabbitMQ, and MongoDB).
  - **mongo:**
    - *check.go:* Implementation of MongoDB check.
    - *check_test.go:* Tests for MongoDB check.
    - *stub.go:* Stub for MongoDB tests.
  - **rabbit:**
    - *check.go:* Implementation of RabbitMQ check.
    - *check_test.go:* Tests for RabbitMQ check.
    - *stub.go:* Stub for RabbitMQ tests.
  - **sqlServer:**
    - *check.go:* Implementation of SQL Server check.
    - *check_test.go:* Tests for SQL Server check.
    - *stub.go:* Stub for SQL Server tests.

- **healthcheck.go:** Main configuration file.

- **docker-compose.yml:** Docker Compose file for tests.

### Service List

1. MongoDB;
2. SQL Server;
3. RabbitMQ.

## Installation

### Using *go get*

```bash
$ go get github.com/wesleycosta/healthcheck-go
```

### Using govendor

```bash
$ govendor add github.com/wesleycosta/healthcheck-go
$ govendor add github.com/wesleycosta/healthcheck-go/checks
$ govendor add github.com/wesleycosta/healthcheck-go/checks/mongo
$ govendor add github.com/wesleycosta/healthcheck-go/checks/rabbit
$ govendor add github.com/wesleycosta/healthcheck-go/checks/sqlServer
```

## Using the library

To configure Healthcheck in your application, follow these steps:

1. **Create a HealthCheck instance:**
   - Use `HealthCheckLib.New()` to create a Healthcheck instance.
   
2. **Configure which services will be monitored:**
   - For each service you want to monitor, configure specific options.
   - Examples:
     - MongoDB: Use the `mongo.Config` file.
     - RabbitMQ: Use the `rabbit.Config` file.
     - SQL Server: Use the `sqlServer.Config` file.

3. **Add Configurations to HealthCheck:**
   - After creating specific service configurations, add them to the HealthCheck instance using the `AddService` method.
   
4. **Integrate HealthCheck into your Endpoint:**
   - When you have a configured HealthCheck instance, integrate it into your endpoint to perform service checks.
   - Return the results of these checks as a JSON response.

#### Configuration Example

##### HealthCheck Configuration File
```go
package healthcheck

import (
	"github.com/gin-gonic/gin"

	HealthCheckLib "github.com/wesleycosta/healthcheck-go"
	"github.com/wesleycosta/healthcheck-go/checks/mongo"
	"github.com/wesleycosta/healthcheck-go/checks/rabbit"
)

func createHealthCheck() HealthCheckLib.HealthCheck {
	mongoConfig := &mongo.Config{
		Url:         config.Get().MongoURL,
		User:        config.Get().MongoUser,
		Password:    config.Get().MongoPassword,
		Database:    config.Get().MongoDatabase,
		AuthSource:  config.Get().MongoAuthSource,
		Timeout:     3,
		ForceTLS:    config.Get().ForceTLS,
		MaxPoolSize: 100,
	}

	rabbitConfig := &rabbit.Config{
		ConnectionString: config.Get().ConnQueue,
	}

	healthCheck := HealthCheckLib.New()
	healthCheck.AddService(mongoConfig)
	healthCheck.AddService(rabbitConfig)

	return healthCheck
}

```

#### Creating an endpoint for HealthCheck

```go
package api

import (
	"github.com/gin-gonic/gin"
	"github.com/wesleycosta/healthcheck"
)

func Base(router *gin.Engine) {
	router.GET("/healthcheck", healthcheck.Endpoint)
}
```

### Healthy Response
```json
{
    "status": "Healthy",
    "results": {
        "mongo": {
            "status": "Healthy",
            "description": "rabbit is healthy"
        },
        "rabbit": {
            "status": "Healthy",
            "description": "rabbit is healthy"
        }
    }
}
```

### Unhealthy Response
```json
{
    "status": "Unhealthy",
    "results": {
        "mongo": {
            "status": "Unhealthy",
            "description": "<error description>"
        },
        "rabbit": {
            "status": "Healthy",
            "description": "rabbit is healthy"
        }
    }
}
```
