package httpclient

import (
	"context"
	"log"
	"os"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.34.0"
)

func InitOTel(serviceName string) func() {
	// Create base resource with service info
	baseRes := resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceName(serviceName),
		semconv.ServiceVersion("1.0.0"),
	)

	envRes, err := resource.Merge(
		resource.Default(),
		resource.Environment(),
	)
	if err != nil {
		log.Fatalf("Failed to merge environment resource: %v", err)
	}

	res, err := resource.Merge(envRes, baseRes)
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
