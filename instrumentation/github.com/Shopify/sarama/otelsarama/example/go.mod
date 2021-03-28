module go.opentelemetry.io/contrib/instrumentation/github.com/Shopify/sarama/otelsarama/example

go 1.14

replace (
	go.opentelemetry.io/contrib => ../../../../../../
	go.opentelemetry.io/contrib/instrumentation/github.com/Shopify/sarama/otelsarama => ../
)

require (
	github.com/DataDog/sketches-go v0.0.1 // indirect
	github.com/Shopify/sarama v1.27.2
	github.com/google/gofuzz v1.1.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/github.com/Shopify/sarama/otelsarama v0.15.1
	go.opentelemetry.io/otel v0.19.0
	go.opentelemetry.io/otel/exporters/stdout v0.19.0
	go.opentelemetry.io/otel/sdk v0.19.0
)
