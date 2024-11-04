FROM golang:1.20

WORKDIR /app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./

RUN update-ca-certificates && go mod download && go mod verify
RUN go install github.com/swaggo/swag/cmd/swag@latest

COPY . .

RUN swag init -g cmd/starter-pack-api/main.go
RUN go build -v -o /usr/local/bin/starter-pack-api cmd/starter-pack-api/main.go

ENTRYPOINT [ "starter-pack-api"]
