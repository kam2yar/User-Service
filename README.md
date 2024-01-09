# User service
Small service for manage the user entity

### Requirements:
1. docker
2. docker-compose

## Run the application
```
make up
```

## Stop the application
```
make down
```

## Watch application logs
```
make logs
```

## Access to container bash
```
make bash
```

## Build the application's docker image
```
make build
```

## Force recreate the app container after build
```
make recreate
```

## Generate protobuf files
Enter to the containers bash using `make bash` then enter:
```
protoc -I ./api --go_out ./api --go_opt paths=source_relative --go-grpc_out ./api --go-grpc_opt paths=source_relative ./api/management.proto
```