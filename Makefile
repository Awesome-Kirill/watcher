lint:
	golangci-lint run --fix

test:
	go test  ./... -v --race --tags=tests

swagger:
	swag init -g .\cmd\main.go
