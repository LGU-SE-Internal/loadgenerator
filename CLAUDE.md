# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a load generator for the Train-Ticket microservices system (41 microservices). It simulates realistic user behavior using a chain-of-responsibility pattern with probabilistic behavior selection. The generator is written in Go and includes OpenTelemetry tracing support.

## Build and Run Commands

### Build
```bash
go build -o loadgenerator
```

### Run with default settings
```bash
./loadgenerator
```

### Run with custom configuration
```bash
# Run with multiple threads and custom sleep duration
./loadgenerator --threads 5 --sleep 1000

# Run a specific behavior chain
./loadgenerator --chain NormalPreserveChain --count 10

# Enable debug logging
./loadgenerator --debug --threads 3
```

### Available command-line flags
- `--threads, -t`: Number of concurrent threads (default: 1)
- `--sleep, -s`: Sleep duration between requests in milliseconds (default: 1000)
- `--chain`: Execute a specific behavior chain by name
- `--count`: Number of times to run the specified chain (default: 1)
- `--debug`: Enable debug logging with caller information
- `--stats`: Show current latency statistics

### Testing
```bash
# Run all tests
go test ./...

# Run tests for a specific package
go test ./behaviors
go test ./service
```

## Code Quality

We use pre-commit hooks to maintain Go code quality. The configuration includes:
- **golangci-lint**: Comprehensive linting with 30+ checks including security, performance, and style
- **gofmt/goimports**: Automatic code formatting and import organization
- **govet**: Static analysis for suspicious constructs
- **Standard hooks**: Trailing whitespace, EOF fixes, file validation

To set up pre-commit:
```bash
# Install pre-commit
pip install pre-commit

# Install git hooks
pre-commit install

# Run on all files
pre-commit run --all-files

# Run specific checks
pre-commit run --all-files --show-diff-on-failure --color=always go-fmt
pre-commit run --all-files --show-diff-on-failure --color=always golangci-lint
```

# Run tests with verbose output
go test -v ./...
```

# Run a specific test
go test -v ./behaviors -run TestSpecificFunction
```

### Docker
```bash
# Build Docker image
docker build -t loadgenerator .

# Run in Docker
docker run loadgenerator --threads 3 --sleep 1000
```

## Environment Variables

- `BASE_URL`: Target service base URL (default: `http://10.10.10.220:30080`)
- `OTEL_EXPORTER_OTLP_ENDPOINT`: OpenTelemetry collector endpoint (default: `http://opentelemetry-collector-deployment.monitoring:4317`)
- `OTEL_EXPORTER_OTLP_HEADERS`: Optional OTLP headers for authentication

## Architecture

### Core Design Pattern: Chain of Responsibility

The load generator uses a chain-of-responsibility pattern where user behaviors are decomposed into reusable nodes that execute sequentially. Chains can have probabilistic next chains, enabling realistic behavior distribution.

**Key structures:**
- `Node`: Atomic operation (login, query tickets, payment, etc.)
- `Chain`: Sequence of nodes representing a complete business flow
- `Context`: Data passing object that carries state between nodes

### Package Structure

- `behaviors/`: Behavior chain definitions and node implementations
  - Contains all user behavior chains (preserve, payment, search, etc.)
  - Defines reusable nodes for common operations
  - Context key constants for data passing between nodes

- `service/`: HTTP service client wrappers
  - `SvcImpl`: Main service client with shared HTTP client
  - Individual service files for each microservice endpoint
  - All services use the shared `httpclient.HttpClient` with OpenTelemetry tracing

- `httpclient/`: HTTP client with OpenTelemetry integration
  - Automatic HTTP request tracing with spans
  - Request/response size tracking
  - Error recording and status tracking

- `stats/`: Latency statistics collection and reporting
  - Tracks request latency percentiles (P50, P95, P99)
  - Identifies slowest endpoints

- `main.go`: Entry point with CLI configuration
  - Defines available behavior chains with probability weights
  - Configures load generator with threads and sleep duration

### Behavior Chains

Available chains (defined in `main.go:13-28`):
- `NormalPreserveChain`: Normal ticket booking (20% probability)
- `NormalOrderPayChain`: Order payment (15%)
- `AdvancedSearchChain`: Advanced search booking (18%)
- `TicketCollectAndEnterStationChain`: Ticket collection and station entry (12%)
- `OrderConsignChain`: Order consignment (8%)
- `ConsignListChain`: Consignment list query (6%)
- `OrderChangeChain`: Order change/rebooking (4%)
- `OrderCancelChain`: Order cancellation (2%)
- `AdminBasicInfoChain`: Admin basic info management (3%)
- `AdminOrderChain`: Admin order management (3%)
- `AdminTravelChain`: Admin travel management (3%)
- `AdminRouteChain`: Admin route management (3%)
- `AdminUserChain`: Admin user management (3%)

### Randomness and Diversity Mechanisms

The load generator implements multi-level randomness:

1. **Behavior-level**: Probabilistic chain selection based on weights
2. **Data-level**: Random selection of routes, contacts, trips, food types
3. **Parameter-level**: Random generation of times, weights, verification codes
4. **Strategy-level**: Multiple search strategies (cheapest, fastest, min stations)
5. **Identity-level**: 20% admin users, 80% normal users

### Context Data Flow

The `Context` object passes data between nodes using string keys. Key context variables are defined in `behaviors/preserve_behavior.go:11-50`:
- User info: `AccountID`, `UserId`, `LoginToken`
- Route info: `RouteID`, `StartStation`, `EndStation`
- Trip info: `TripID`, `TrainTypeName`, `DepartureTime`
- Order info: `OrderId`, `Price`, `SeatClass`
- Contact info: `ContactsID`, `Name`, `PhoneNumber`
- Food info: `FoodType`, `FoodName`, `StoreName`
- Insurance info: `AssuranceTypeIndex`, `AssuranceTypeName`

### Load Generator Lifecycle

1. Initialize shared `SvcImpl` client with OpenTelemetry
2. Spawn N worker goroutines (configured by `--threads`)
3. Each worker executes the composed chain in a loop
4. Sleep between iterations (configured by `--sleep`)
5. Graceful shutdown on SIGINT/SIGTERM
6. Statistics printed every 10 seconds
7. Old statistics cleaned every 5 minutes
8. Garbage collection every 2 minutes

## Development Guidelines

### Adding a New Behavior Chain

1. Define nodes in `behaviors/` (e.g., `new_behavior.go`)
2. Create a chain by composing nodes: `NewChain(node1, node2, ...)`
3. Register the chain in `main.go` chains map
4. Add to composed chain with probability weight in `main.go:74-87`

### Adding a New Service Client

1. Create service file in `service/` (e.g., `new_service.go`)
2. Add methods to `SvcImpl` that use `s.cli.SendRequest()` or `s.cli.SendRequestWithContext()`
3. Use `s.BaseUrl` for constructing endpoint URLs
4. Return structured response types defined in `service/utils.go`

### Context Usage Pattern

```go
// Set data in context
ctx.Set(KeyName, value)

// Get data from context
value := ctx.Get(KeyName).(ExpectedType)

// Get shared client
client := ctx.Get(Client).(*service.SvcImpl)
```

### Node Implementation Pattern

```go
func MyNode(ctx *Context) (*NodeResult, error) {
    // Get dependencies from context
    client := ctx.Get(Client).(*service.SvcImpl)

    // Perform operation
    result, err := client.SomeService()
    if err != nil {
        return nil, err
    }

    // Store results in context for next nodes
    ctx.Set(SomeKey, result.Data)

    return nil, nil
}
```

### Service Client Pattern

All service methods follow this pattern:
```go
func (s *SvcImpl) ServiceMethod(input *RequestType) (*ResponseType, error) {
    url := s.BaseUrl + "/api/v1/endpoint"
    resp, err := s.cli.SendRequest("POST", url, input)
    if err != nil {
        return nil, err
    }

    var result ResponseType
    json.Unmarshal(resp, &result)
    return &result, nil
}
```

## OpenTelemetry Tracing

All HTTP requests are automatically traced. Traces include:
- HTTP method, URL, status code
- Request/response content length
- Error recording for failed requests
- Span events for request logging

Configure the OTLP endpoint via environment variable to send traces to Jaeger, Grafana, or other collectors.

## Important Implementation Notes

- The shared `SvcImpl` client is reused across all worker goroutines for efficiency
- Each worker creates a new `Context` for each chain execution
- Nodes should be stateless and only communicate via `Context`
- All randomness uses `math/rand` for reproducibility
- Graceful shutdown waits up to 30 seconds for workers to complete
- Statistics are thread-safe and collected globally via `stats.GlobalLatencyManager`
