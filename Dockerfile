# Docker image for the Drone Artifactory plugin
#
#   docker build -t jmccann/drone-artifactory .

#
# Build golang binary
#

FROM golang:1.12-alpine3.9 AS builder

RUN apk add --no-cache build-base ca-certificates git

RUN mkdir -p /tmp/drone-artifactory
WORKDIR /tmp/drone-artifactory

COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download

COPY . .

RUN go test ./...
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -o /go/bin/drone-artifactory

#
# Build docker image
#

FROM alpine:3.9

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /go/bin/drone-artifactory /bin/
ENTRYPOINT ["/bin/drone-artifactory"]
