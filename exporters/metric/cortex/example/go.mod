module go.opentelemetry.io/contrib/exporters/metric/cortex/example

go 1.15

replace (
	go.opentelemetry.io/contrib/exporters/metric/cortex => ../
	go.opentelemetry.io/contrib/exporters/metric/cortex/utils => ../utils/
)

require (
	go.opentelemetry.io/contrib/exporters/metric/cortex v0.25.0
	go.opentelemetry.io/contrib/exporters/metric/cortex/utils v0.25.0
	go.opentelemetry.io/otel v1.4.0
	go.opentelemetry.io/otel/metric v0.27.0
	go.opentelemetry.io/otel/sdk v1.4.0
	go.opentelemetry.io/otel/sdk/metric v0.27.0
)
