module github.com/harsinigovindarao/braindump-cli

//module github.com/harsinigovindarao/braindump-cli

go 1.24.4

require (
	github.com/google/uuid v1.6.0
	google.golang.org/grpc v1.74.2
	google.golang.org/protobuf v1.36.7
)

require (
	golang.org/x/net v0.40.0 // indirect
	golang.org/x/sys v0.33.0 // indirect
	golang.org/x/text v0.25.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250528174236-200df99c418a // indirect
)

replace github.com/harsinigovindarao/braindump-cli => ./braindump-cli

replace github.com/harsinigovindarao/braindump-cli/internal/utils => ./internal/utils

replace github.com/harsinigovindarao/braindump-cli/internal/nlp/proto => ./internal/nlp/proto
