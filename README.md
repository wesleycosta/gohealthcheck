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
$ govendor add github.com/wesleycosta/healthcheck-go/checks/sqlServer
```

## Utilizando a lib

Para configurar o Healthcheck na sua aplicação, siga os seguintes passos:

1. **Crie uma instância do HealthCheck:**
   - Utilize `HealthCheckLib.New()` para criar uma instância do Healthcheck.
   
2. **Configure quais serão os serviços monitorados:**
   - Para cada serviço que deseja monitorar, configure as opções específicas.
   - Exemplos:
     - MongoDB: Utilize o arquivo `mongo.Config`.
     - RabbitMQ: Utilize o arquivo `rabbit.Config`.
     - SQL Server: Utilize o arquivo `sqlServer.Config`.

3. **Adicione as Configurações ao HealthCheck:**
   - Após criar as configurações específicas do serviço, adicione-as à instância do HealthCheck utilizando o método `AddService`.
   
4. **Integre o HealthCheck em seu Endpoint:**
   - Ao criar uma instância do HealthCheck configurada, integre-o em seu endpoint para executar verificações dos serviços.
   - Retorne o resultado dessas verificações como uma resposta JSON.


#### Exemplo do configuração

##### Arquivo de configuração do HealthCheck
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

#### Criando um endpoint para o HealthCheck

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
