FROM docker.bintray.io/jfrog/jfrog-cli-go:1.24.0

ADD drone-artifactory /bin/
ENTRYPOINT ["/bin/drone-artifactory"]