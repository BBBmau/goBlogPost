# syntax=docker/dockerfile:1
FROM golang:1.16-alpine
WORKDIR /BlogWebsite
COPY . .
RUN go build

CMD ["./v1"]

