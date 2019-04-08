# Docker image for the Drone Artifactory plugin
#
#     docker build -t jmccann/drone-artifactory .

#
# Get CLI
#

FROM docker.bintray.io/jfrog/jfrog-cli-go:1.24.0 AS cli

#
# Build golang binary
#

FROM golang:1.12 AS builder

COPY --from=cli /usr/local/bin/jfrog /bin/jfrog

RUN mkdir -p /tmp/drone-artifactory
WORKDIR /tmp/drone-artifactory
COPY . .

RUN go mod download
RUN go test -tags cli
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -o /go/bin/drone-artifactory

#
# Build docker image
#

FROM ubuntu:18.04

RUN apt update && apt install -y ca-certificates && apt clean && rm -rf /var/lib/apt/lists/*

COPY --from=cli /usr/local/bin/jfrog /bin/jfrog
COPY --from=builder /go/bin/drone-artifactory /bin/
ENTRYPOINT ["/bin/drone-artifactory"]
