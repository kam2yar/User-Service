# User service
Small service for manage the user entity

### Requirements:
1. docker
2. docker-compose

## Run the containers
```
make up
```

## Stop the containers
```
make down
```

## Watch application logs
```
make logs
```

## Access to the app container bash
```
make bash
```

## Build executable app
```
make executable
```

## Build the docker image for app
```
make build
```

## Force recreate the app container after build
```
make recreate
```

## Restart docker containers
```
make restart
```

## Generate protobuf files
Enter to the containers bash using `make bash` then enter:
```
protoc -I ./api --go_out ./api --go_opt paths=source_relative --go-grpc_out ./api --go-grpc_opt paths=source_relative ./api/management.proto
```