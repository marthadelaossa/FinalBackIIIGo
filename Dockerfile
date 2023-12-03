FROM golang:1.21.4 AS build

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /usr/local/bin/app ./cmd/...

# Imagen final
FROM alpine:latest

WORKDIR /root/

COPY --from=build /usr/local/bin/app .

EXPOSE 8080

CMD ["./app"]