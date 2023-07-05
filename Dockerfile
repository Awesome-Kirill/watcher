FROM golang:1.20

WORKDIR /usr/src/app

COPY go.mod go.sum site.txt ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -ldflags="-s -w" -o /usr/local/bin/app /usr/src/app/cmd/main.go

CMD ["app"]

