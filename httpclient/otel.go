package httpclient

import (
	"context"
	"log"
	"os"
	"strings"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.34.0"
)

func InitOTel(serviceName string) func() {
	// Base attributes
	attrs := []attribute.KeyValue{
		semconv.ServiceName(serviceName),
		semconv.ServiceVersion("1.0.0"),
	}

	// Add custom attributes from environment variables
	if serviceNamespace := os.Getenv("SERVICE_NAMESPACE"); serviceNamespace != "" {
		attrs = append(attrs, attribute.String("service.namespace", serviceNamespace))
	}
	
	if customServiceName := os.Getenv("SERVICE_NAME"); customServiceName != "" {
		// Override service name if provided via environment
		attrs[0] = semconv.ServiceName(customServiceName)
	}
	
	if podName := os.Getenv("POD_NAME"); podName != "" {
		attrs = append(attrs, attribute.String("pod.name", podName))
	}

	// Parse OTEL_RESOURCE_ATTRIBUTES if present
	if otelResourceAttrs := os.Getenv("OTEL_RESOURCE_ATTRIBUTES"); otelResourceAttrs != "" {
		for _, pair := range strings.Split(otelResourceAttrs, ",") {
			if kv := strings.SplitN(strings.TrimSpace(pair), "=", 2); len(kv) == 2 {
				key := strings.TrimSpace(kv[0])
				value := strings.TrimSpace(kv[1])
				
				// Handle standard service attributes
				switch key {
				case "service.name":
					attrs[0] = semconv.ServiceName(value)
				case "service.namespace":
					attrs = append(attrs, attribute.String("service.namespace", value))
				case "pod.name":
					attrs = append(attrs, attribute.String("pod.name", value))
				default:
					attrs = append(attrs, attribute.String(key, value))
				}
			}
		}
	}

	res, err := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			attrs...,
		),
	)
	if err != nil {
		log.Fatalf("Failed to create resource: %v", err)
	}

	var exporter trace.SpanExporter

	otlpEndpoint := os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT")
	if otlpEndpoint == "" {
		otlpEndpoint = "opentelemetry-collector-deployment.monitoring:4317"
		log.Printf("No OTEL_EXPORTER_OTLP_ENDPOINT set, using default: %s", otlpEndpoint)
	}

	log.Printf("Using OTLP gRPC exporter with endpoint: %s", otlpEndpoint)

	otlpExporter, err := otlptracegrpc.New(
		context.Background(),
		otlptracegrpc.WithEndpoint(otlpEndpoint),
		otlptracegrpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalf("Failed to create OTLP exporter: %v", err)
	}
	exporter = otlpExporter

	tp := trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithResource(res),
		trace.WithSampler(trace.AlwaysSample()),
	)

	otel.SetTracerProvider(tp)

	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	))

	return func() {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := tp.Shutdown(ctx); err != nil {
			log.Printf("Failed to shutdown trace provider: %v", err)
		}
	}
}
