# Этап, на котором выполняется сборка приложения
FROM golang:1.18-alpine as builder
WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN go build -v -ldflags="-s -w" -o /app cmd/main.go


FROM alpine:3.15
WORKDIR /usr/src/app
COPY --from=builder app /usr/local/bin/app
COPY site.txt .
ENTRYPOINT ["app"]
