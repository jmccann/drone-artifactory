# drone-artifactory

Drone plugin to publish artifacts from the build to [Artifactory](https://www.jfrog.com/artifactory/).

## Build

Build the binary with the following command:

```sh
go build
```

## Test

Test the code with the following command:

```sh
go test -v -race ./...
```

## Docker

Build the docker image with the following command:

```sh
docker build -t jmccann/drone-artifactory .
```
