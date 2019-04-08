# Docker image for the Drone Artifactory plugin
#
#     docker build -t jmccann/drone-artifactory .

#
# Builder
#

FROM golang:1.12 AS builder

ENV CLI_VERSION 1.19.0

RUN apt update && apt install -y ca-certificates && apt clean && rm -rf /var/lib/apt/lists/*
ADD https://jfrog.bintray.com/jfrog-cli-go/${CLI_VERSION}/jfrog-cli-linux-amd64/jfrog /bin/jfrog
RUN chmod 0755 /bin/jfrog

RUN mkdir -p /tmp/drone-artifactory
WORKDIR /tmp/drone-artifactory
COPY . .

RUN go mod download
RUN go test
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -o /go/bin/drone-artifactory

#
# Published Image
#

FROM ubuntu:18.04

RUN apt update && apt install -y ca-certificates && apt clean && rm -rf /var/lib/apt/lists/*

COPY --from=builder /bin/jfrog /bin/jfrog
COPY --from=builder /go/bin/drone-artifactory /bin/
ENTRYPOINT ["/bin/drone-artifactory"]
