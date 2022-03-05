# Transaction service

This holds the public api for the transaction service.
APIs can be viewed in the swagger docs at http://localhost:8080/swagger/

Architecture: Clean code architecture and dependency injection.

This service makes a client connection to the wallet grpc service.

### Folders and files

1. .env: replace the values with yours or use the default values

2. pkg: contains system wide configurations and extensions

3. model: contains necessary structs for the domain layer

4. repository: holds the apps repository

5. service: contains core implementation of logic

### Run all tests

go test -v -count=1 ./...

### Build

go build

### Run

go run .
