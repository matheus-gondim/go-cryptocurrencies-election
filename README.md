# Cryptocurrencies Election

## Generate the protos

    protoc --go_out=. --go-grpc_out=. ./proto/*.proto

## Run migrations

    go run migration/migration.go

## Test coverage generator

    mkdir coverage
    go test ./... --coverprofile ./coverage/cover.out
    go tool cover -html=./coverage/cover.out -o ./coverage/cover.html