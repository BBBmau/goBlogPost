# syntax=docker/dockerfile:1
FROM golang:1.16-alpine
WORKDIR /BlogWebsite
COPY goWebServers.go go.mod go.sum ./
RUN go mod download

CMD ["go", "run", "goWebServers.go"]

