package observability

import (
	"context"
	"errors"
	"time"

	"github.com/felipeversiane/go-otel/internal/infra/config"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
)

type observer struct {
	config config.ObservabilityConfig
}

type ObserverInterface interface {
	SetupOtel(context.Context) (shutdown func(context.Context) error, err error)
}

func NewObserver(config config.ObservabilityConfig) ObserverInterface {
	return observer{config}
}

func (o observer) SetupOtel(ctx context.Context) (shutdown func(context.Context) error, err error) {
	var shutdownFuncs []func(context.Context) error

	// shutdown calls cleanup functions registered via shutdownFuncs.
	// The errors from the calls are joined.
	// Each registered cleanup will be invoked once.
	shutdown = func(ctx context.Context) error {
		var err error
		for _, fn := range shutdownFuncs {
			err = errors.Join(err, fn(ctx))
		}
		shutdownFuncs = nil
		return err
	}

	// handleErr calls shutdown for cleanup and makes sure that all errors are returned.
	handleErr := func(inErr error) {
		err = errors.Join(inErr, shutdown(ctx))
	}

	// Setup resource.
	res, err := newResource(o.config)
	if err != nil {
		handleErr(err)
		return
	}

	// Setup trace provider.
	tracerProvider, err := newTraceProvider(ctx, o.config, res)
	if err != nil {
		handleErr(err)
		return
	}
	shutdownFuncs = append(shutdownFuncs, tracerProvider.Shutdown)
	otel.SetTracerProvider(tracerProvider)

	// Setup meter provider.
	meterProvider, err := newMeterProvider(ctx, o.config, res)
	if err != nil {
		handleErr(err)
		return
	}
	shutdownFuncs = append(shutdownFuncs, meterProvider.Shutdown)
	otel.SetMeterProvider(meterProvider)

	return shutdown, err
}

func newResource(config config.ObservabilityConfig) (*resource.Resource, error) {
	return resource.Merge(resource.Default(),
		resource.NewWithAttributes(semconv.SchemaURL,
			semconv.ServiceName(config.ServiceName),
			semconv.ServiceVersion(config.ServiceVersion),
		))
}

func newTraceProvider(ctx context.Context, config config.ObservabilityConfig, res *resource.Resource) (*trace.TracerProvider, error) {
	options := []otlptracegrpc.Option{}

	if config.OtelExporterOtlpEndpoint != "" {
		options = append(options, otlptracegrpc.WithEndpoint(config.OtelExporterOtlpEndpoint))
	}

	if config.OtelExporterOtlpInsecure {
		options = append(options, otlptracegrpc.WithInsecure())
	}

	traceExporter, err := otlptracegrpc.New(ctx, options...)
	if err != nil {
		return nil, err
	}

	traceProvider := trace.NewTracerProvider(
		trace.WithBatcher(traceExporter,
			// Default is 5s. Set to 1s for demonstrative purposes.
			trace.WithBatchTimeout(time.Second)),
		trace.WithResource(res),
	)
	return traceProvider, nil
}

func newMeterProvider(ctx context.Context, config config.ObservabilityConfig, res *resource.Resource) (*metric.MeterProvider, error) {
	options := []otlpmetricgrpc.Option{}

	if config.OtelExporterOtlpEndpoint != "" {
		options = append(options, otlpmetricgrpc.WithEndpoint(config.OtelExporterOtlpEndpoint))
	}

	if config.OtelExporterOtlpInsecure {
		options = append(options, otlpmetricgrpc.WithInsecure())
	}
	metricExp, err := otlpmetricgrpc.New(ctx, options...)
	if err != nil {
		return nil, err
	}

	meterProvider := metric.NewMeterProvider(
		metric.WithResource(res),
		metric.WithReader(metric.NewPeriodicReader(metricExp,
			// Default is 1m. Set to 3s for demonstrative purposes.
			metric.WithInterval(3*time.Second))),
	)
	otel.SetMeterProvider(meterProvider)
	return meterProvider, nil
}
