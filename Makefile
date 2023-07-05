lint:
	golangci-lint run --fix

test:
	go test  ./... -v --race

swagger:
	swag init -g .\cmd\main.go
