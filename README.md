# User service
Small service for manage the user entity

### Requirements:
1. docker
2. docker-compose

# Installation

# Automatic
You can setup applications with:
```
make setup
```

# Manual
You can also do any step by yourself

## Copy .env.example to .env
```
make env
```

## Run the containers
```
make up
```
## Migrate database
```
make migrate
```
## Build executable app
```
make executable
```
## Generate protobuf files
```
make pb
```

# Additional commands
These are some important commands that you can use after setup the application

## Watch application logs
```
make logs
```

## Access to the app container bash
```
make bash
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