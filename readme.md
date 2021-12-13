# Reproduce gRPC Reflection Failure

1. Run the server with `go run main.go`
2. List services with `grpcurl -plaintext localhost:8080 list` - this works
3. Try to describe `simple.SimpleService` with `grpcurl -plaintext localhost:8080 describe simple.SimpleService`

Receive the following error:

```
Failed to resolve symbol "simple.SimpleService": Symbol not found: simple.SimpleService
caused by: File not found: github.com/danielvladco/go-proto-gql/pb/graphql.proto
```