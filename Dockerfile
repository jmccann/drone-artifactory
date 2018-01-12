# Docker image for the Drone Artifactory plugin
#
#     docker build -t jmccann/drone-artifactory .

FROM ubuntu:17.04

ENV CLI_VERSION 1.13.1
ADD https://jfrog.bintray.com/jfrog-cli-go/${CLI_VERSION}/jfrog-cli-linux-amd64/jfrog /bin/jfrog
RUN apt update && apt install -y ca-certificates && apt clean && rm -rf /var/lib/apt/lists/*

RUN chmod 0755 /bin/jfrog

ADD drone-artifactory /bin/
ENTRYPOINT ["/bin/drone-artifactory"]
