# User service
Small service for manage the user entity

### Requirements:
1. docker
2. docker-compose

## Copy .env.example to .env

## Run the containers
```
make up
```
## Migrate database
```
make migrate
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

## Stop the containers
```
make down
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
```
make pb
```