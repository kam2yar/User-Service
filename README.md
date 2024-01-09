# User service
Small service for manage the user entity

# Generate protobuf files
```
protoc -I ./api --go_out ./api --go_opt paths=source_relative --go-grpc_out ./api --go-grpc_opt paths=source_relative ./api/management.proto
```