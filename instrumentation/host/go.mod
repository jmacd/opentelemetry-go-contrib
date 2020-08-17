module go.opentelemetry.io/contrib/instrumentation/host

go 1.14

replace go.opentelemetry.io/contrib => ../..

require (
	github.com/shirou/gopsutil v2.20.7+incompatible
	github.com/stretchr/testify v1.6.1
	go.opentelemetry.io/contrib v0.10.1
	go.opentelemetry.io/otel v0.10.0
	go.opentelemetry.io/otel/exporters/stdout v0.10.0
	go.opentelemetry.io/otel/sdk v0.10.0
)
