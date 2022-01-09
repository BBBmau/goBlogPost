# syntax=docker/dockerfile:1
FROM golang:1.16-alpine
WORKDIR /BlogWebsite
<<<<<<< HEAD
COPY goWebServers.go go.mod go.sum ./
RUN go mod download

CMD ["go", "run", "goWebServers.go"]
=======
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY *.go ./
RUN go build -o /docker-gs-ping
CMD ["/docker-gs-ping"]
>>>>>>> 33c414a73bba58bba45d504913c47059e50767d4

