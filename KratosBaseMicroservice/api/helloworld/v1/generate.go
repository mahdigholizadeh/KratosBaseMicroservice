package v1

// Note: Ensure the path to third_party is correct (e.g., ../../../third_party)

//go:generate protoc --proto_path=. --proto_path=../../../third_party --go_out=. --go_opt=paths=source_relative --go-http_out=. --go-http_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative greeter.proto
