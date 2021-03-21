module go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp/example

go 1.14

replace (
	go.opentelemetry.io/contrib => ../../../../../
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp => ../
)

require (
	github.com/DataDog/sketches-go v0.0.1 // indirect
	github.com/google/gofuzz v1.1.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.15.1
	go.opentelemetry.io/otel v0.19.0
	go.opentelemetry.io/otel/exporters/stdout v0.19.0
	go.opentelemetry.io/otel/sdk v0.19.0
)
