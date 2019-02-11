# drone-artifactory

[![Build Status](https://cloud.drone.io/api/badges/jmccann/drone-artifactory/status.svg)](https://cloud.drone.io/jmccann/drone-artifactory)

Drone plugin to publish artifacts from the build to [Artifactory](https://www.jfrog.com/artifactory/).

## Build & Test

Build/test the binary with the following commands:

```sh
export GO111MODULE=on
go test -cover ./...
go build
```

## Docker

[Drone CLI](http://docs.drone.io/cli-installation/) is required.

Build the docker image with the following commands:

```sh
drone exec
```

## Usage

Execute from the working directory:

```sh
docker run --rm \
  jmccann/drone-artifactory --url https://myarti.com/artifactory \
  --username JohnDoe --password abcd1234 \
  --source *.go --path repo-key/path/to/target
```
