FROM docker.bintray.io/jfrog/jfrog-cli-go:1.24.0

FROM scratch

COPY --from=0 /usr/local/bin/jfrog /bin/jfrog
ADD drone-artifactory /bin/
ENTRYPOINT ["/bin/drone-artifactory"]
