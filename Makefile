lint:
	golangci-lint run

test:
	go test  ./... -v --race

swagger:
	swag init -g .\cmd\main.go
