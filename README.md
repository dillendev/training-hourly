# Hourly

## REST API

The Hourly API is a REST API and includes an [OpenAPI specification](https://app.swaggerhub.com/apis-docs/dillen.dev/hourly-api/0.0.1) . It isn't published on the internet but can be run locally by using the following command:

```bash
go run github.com/dillendev/training-hourly/cmd/server@latest
```

This will run the Hourly API on port 8989 on localhost and contains the actual data for the assignment.

## Go client

The Go client is a client for the Hourly API. It can be used to retrieve the data from the API and can be used to create a CLI or a web application. Note that it requires authorization by using a token. The token can be retrieved by using the `/api/auth/tokens` endpoint.

### Usage

```go
package main

import (
	"github.com/deepmap/oapi-codegen/pkg/securityprovider"
	hourly "github.com/dillendev/training-hourly"
)

func main() {
    // Setup bearer token authentication
	provider, err := securityprovider.NewSecurityProviderBearerToken("my-token")
	if err != nil {
		panic(err)
	}

	// Setup client-side middleware to add the authorization header
    authOption := hourly.WithRequestEditorFn(provider.Intercept)

	// Create a new client for the Hourly API with the middleware
	client, err := hourly.NewClientWithResponses("http://localhost:8989", authOption)
	if err != nil {
		panic(err)
	}

	// Do something with: client
}
```
