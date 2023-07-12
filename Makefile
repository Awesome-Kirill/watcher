lint:
	golangci-lint run --fix

test:
	go test  ./... -v --race --count=10 --tags=tests

swagger:
	swag init -g .\cmd\main.go
