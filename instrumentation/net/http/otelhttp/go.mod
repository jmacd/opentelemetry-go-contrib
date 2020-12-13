module go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp

go 1.14

replace go.opentelemetry.io/contrib => ../../../..

require (
	github.com/felixge/httpsnoop v1.0.1
	github.com/stretchr/testify v1.6.1
	go.opentelemetry.io/contrib v0.15.0 // indirect
	go.opentelemetry.io/otel v0.12.0
)
