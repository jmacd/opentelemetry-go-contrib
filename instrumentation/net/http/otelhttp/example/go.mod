module go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp/example

go 1.18

replace go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp => ../

require (
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.36.4
	go.opentelemetry.io/otel v1.14.0
	go.opentelemetry.io/otel/exporters/stdout/stdoutmetric v0.33.0
	go.opentelemetry.io/otel/exporters/stdout/stdouttrace v1.11.1
	go.opentelemetry.io/otel/metric v0.33.0
	go.opentelemetry.io/otel/sdk v1.14.0
	go.opentelemetry.io/otel/sdk/metric v0.33.0
	go.opentelemetry.io/otel/trace v1.14.0
)

require (
	github.com/felixge/httpsnoop v1.0.3 // indirect
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	golang.org/x/sys v0.5.0 // indirect
)
