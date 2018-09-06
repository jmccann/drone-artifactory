# drone-artifactory

[![Build Status](http://beta.drone.io/api/badges/jmccann/drone-artifactory/status.svg)](http://beta.drone.io/jmccann/drone-artifactory)

Drone plugin to publish artifacts from the build to [Artifactory](https://www.jfrog.com/artifactory/).

## Build

Build the binary with the following commands:

```sh
go build
go test
```

## Docker

[Drone CLI](http://docs.drone.io/cli-installation/) is required.

Build the docker image with the following commands:

```sh
drone exec
docker build -t jmccann/drone-artifactory .
```

## Usage

Execute from the working directory:

```sh
docker run --rm \
  jmccann/drone-artifactory --url https://myarti.com/artifactory \
  --username JohnDoe --password abcd1234 \
  --source *.go --path repo-key/path/to/target
```
