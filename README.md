[![GoDoc](https://godoc.org/github.com/wesleycosta/goseq?status.svg)](https://godoc.org/github.com/wesleycosta/goseq)
# Health Check Go

Biblioteca de Health Check em Golang.

## Estrutura do Projeto

- **checks:** Pasta contendo verificações para dependências externas (SQL Server, RabbitMQ e MongoDB).
  - **mongo:**
    - *check.go:* Implementação da verificação para MongoDB.
    - *check_test.go:* Testes para a verificação do MongoDB.
    - *stub.go:* Stub para os testes do MongoDB.
  - **rabbit:**
    - *check.go:* Implementação da verificação para RabbitMQ.
    - *check_test.go:* Testes para a verificação do RabbitMQ.
    - *stub.go:* Stub para os testes do RabbitMQ.
  - **sqlServer:**
    - *check.go:* Implementação da verificação para SQL Server.
    - *check_test.go:* Testes para a verificação do SQL Server.
    - *stub.go:* Stub para os testes do SQL Server.

- **healthcheck.go:** Arquivo principal de configuração.

- **docker-compose.yml:** Arquivo Docker Compose para os testes.

### Lista de serviços

1. MongoDB;
2. SQL Server;
3. RabbitMQ.


## Instalação

### Utilizando *go get*

```bash
$ go get github.com/wesleycosta/healthcheck-go
```

### Utilizando govendor

```bash
$ govendor add github.com/wesleycosta/healthcheck-go
$ govendor add github.com/wesleycosta/healthcheck-go/checks
$ govendor add github.com/wesleycosta/healthcheck-go/checks/mongo
$ govendor add github.com/wesleycosta/healthcheck-go/checks/rabbit
```

## Exemplo de Uso

### HealthCheck

```go
package healthcheck

import (
	"github.com/gin-gonic/gin"
	"github.com/wesleycosta/boleto-api/config"

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

func Endpoint(c *gin.Context) {
	healthcheck := createHealthCheck()
	c.JSON(200, healthcheck.Execute())
}
```

### API

```go
package api

import (
	"github.com/gin-gonic/gin"
	"github.com/wesleycosta/boleto-api/healthcheck"
)

func Base(router *gin.Engine) {
	router.GET("/healthcheck", healthcheck.Endpoint)
}
```
### Resposta Healthy
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

### Resposta Unhealthy
```json
{
    "status": "Unhealthy",
    "results": {
        "mongo": {
            "status": "Unhealthy",
            "description": "<descricao do erro>"
        },
        "rabbit": {
            "status": "Healthy",
            "description": "rabbit is healthy"
        }
    }
}
```
