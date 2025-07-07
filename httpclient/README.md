# HTTP Client with OpenTelemetry Tracing

这个 HTTP 客户端已经集成了 OpenTelemetry 分布式追踪功能，可以自动记录 HTTP 请求的详细信息。

## 功能特性

- 🔍 **自动 HTTP 追踪**: 使用 otelhttp 自动记录 HTTP 请求的 spans
- 📊 **详细监控**: 记录请求方法、URL、状态码、请求/响应大小等
- 🚨 **错误记录**: 自动记录和标记错误状态
- 📈 **统计信息**: 内置请求成功/失败统计
- 🔧 **灵活配置**: 支持自定义 context 传递

## 使用方法

### 1. 初始化 OpenTelemetry

```go
import "github.com/Lincyaw/loadgenerator/httpclient"

// 初始化 OpenTelemetry
cleanup := httpclient.InitOTel("your-service-name")
defer cleanup()
```

### 2. 创建 HTTP 客户端

```go
// 创建带有 tracing 的 HTTP 客户端
client := httpclient.NewCustomClient()

// 添加通用头信息
client.AddHeader("Content-Type", "application/json")
client.AddHeader("Authorization", "Bearer your-token")
```

### 3. 发送请求

```go
// 方式1: 使用默认 context
response, err := client.SendRequest("POST", "https://api.example.com/users", requestBody)

// 方式2: 使用自定义 context（推荐）
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()
response, err := client.SendRequestWithContext(ctx, "GET", "https://api.example.com/users/123", nil)
```

### 4. 获取统计信息

```go
// 获取请求计数
count := client.GetRequestCount()

// 获取详细统计信息
stats := client.GetRequestStats()

// 生成 Markdown 格式的统计表格
table := httpclient.GenerateMarkdownTable(stats)
fmt.Println(table)
```

## Trace 信息

每个 HTTP 请求都会自动创建一个 span，包含以下信息：

### Span 属性
- `http.method`: HTTP 方法 (GET, POST, etc.)
- `http.url`: 请求 URL
- `http.status_code`: 响应状态码
- `http.status_text`: 响应状态文本
- `http.request_content_length`: 请求体大小
- `http.response_content_length`: 响应体大小

### Span 事件
- `request_logged`: 请求记录事件，包含请求详细信息

### Span 状态
- 成功请求 (2xx): `codes.Ok`
- 客户端/服务器错误 (4xx/5xx): `codes.Error`
- 网络错误: `codes.Error` + 错误记录

## 运行示例

```bash
# 编译示例
go build -o examples/otel_example examples/otel_example.go

# 运行示例
./examples/otel_example
```

示例会发送几个 HTTP 请求并将 trace 信息输出到控制台。

## 配置 OTLP Endpoint

### 环境变量配置

你可以通过以下环境变量来配置 OpenTelemetry 导出器：

#### OTEL_EXPORTER_OTLP_ENDPOINT
设置 OTLP endpoint URL。如果不设置此变量，将使用默认值：`http://opentelemetry-collector-deployment.monitoring:4317`

```bash
# 使用 Jaeger (OTLP HTTP)
export OTEL_EXPORTER_OTLP_ENDPOINT=http://localhost:4318

# 使用 OTEL Collector
export OTEL_EXPORTER_OTLP_ENDPOINT=http://otel-collector:4318

# 使用云服务 (例如: Grafana Cloud, DataDog, etc.)
export OTEL_EXPORTER_OTLP_ENDPOINT=https://your-cloud-provider.com/v1/traces
```

#### OTEL_EXPORTER_OTLP_HEADERS (可选)
设置 OTLP 请求头，用于认证等。

```bash
# 使用 API Key 认证
export OTEL_EXPORTER_OTLP_HEADERS="authorization=Bearer your-api-key"

# 使用多个头信息
export OTEL_EXPORTER_OTLP_HEADERS="authorization=Bearer token,x-custom-header=value"
```

### 使用示例

#### 1. 使用 Jaeger

```bash
# 启动 Jaeger (使用 Docker)
docker run -d --name jaeger \
  -p 16686:16686 \
  -p 4318:4318 \
  jaegertracing/all-in-one:latest

# 配置环境变量
export OTEL_EXPORTER_OTLP_ENDPOINT=http://localhost:4318
export BASE_URL=http://your-service:8080

# 运行负载生成器
./loadgenerator --threads 3 --sleep 1000
```

#### 2. 使用 OpenTelemetry Collector

```bash
# 配置 OTEL Collector endpoint
export OTEL_EXPORTER_OTLP_ENDPOINT=http://otel-collector:4318
export BASE_URL=http://your-service:8080

# 运行负载生成器
./loadgenerator --chain NormalPreserveChain --count 10
```

#### 3. 使用云服务

```bash
# 配置云服务 endpoint 和认证
export OTEL_EXPORTER_OTLP_ENDPOINT=https://api.your-cloud-provider.com/v1/traces
export OTEL_EXPORTER_OTLP_HEADERS="authorization=Bearer your-api-token"
export BASE_URL=http://your-service:8080

# 运行负载生成器
./loadgenerator --debug --threads 5
```

### 验证配置

当程序启动时，你会看到以下日志信息之一：

```
# 使用 OTLP exporter
Using OTLP exporter with endpoint: http://localhost:4318

# 使用 stdout exporter（默认）
Using stdout exporter (set OTEL_EXPORTER_OTLP_ENDPOINT to use OTLP)
```

## 集成其他 Exporter

默认使用 stdout exporter 将 trace 输出到控制台。你可以修改 `httpclient/otel.go` 文件来集成其他 exporter，如：

- **Jaeger**: `go.opentelemetry.io/otel/exporters/jaeger`
- **OTLP**: `go.opentelemetry.io/otel/exporters/otlp/otlptrace`
- **Zipkin**: `go.opentelemetry.io/otel/exporters/zipkin`

## 依赖项

项目已自动添加了以下 OpenTelemetry 依赖：

```
go.opentelemetry.io/otel v1.24.0
go.opentelemetry.io/otel/trace v1.24.0
go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.49.0
go.opentelemetry.io/otel/exporters/stdout/stdouttrace v1.24.0
go.opentelemetry.io/otel/sdk v1.24.0
```
