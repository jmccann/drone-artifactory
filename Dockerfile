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

COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download

COPY . .

RUN go test -tags cli
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -o /go/bin/drone-artifactory

#
# Build docker image
#

FROM scratch

COPY --from=cli /usr/local/bin/jfrog /bin/jfrog
COPY --from=builder /go/bin/drone-artifactory /bin/
ENTRYPOINT ["/bin/drone-artifactory"]
