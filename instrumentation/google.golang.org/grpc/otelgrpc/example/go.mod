module go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc/example

go 1.14

replace (
	go.opentelemetry.io/contrib => ../../../../../
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc => ../
)

require (
	github.com/DataDog/sketches-go v0.0.1 // indirect
	github.com/golang/protobuf v1.4.3
	github.com/google/gofuzz v1.1.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.15.1
	go.opentelemetry.io/otel v0.18.0
	go.opentelemetry.io/otel/exporters/stdout v0.18.0
	go.opentelemetry.io/otel/sdk v0.18.0
	golang.org/x/net v0.0.0-20200707034311-ab3426394381
	google.golang.org/grpc v1.34.0
)
