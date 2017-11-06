# Docker image for the Drone Artifactory plugin
#
#     docker build -t jmccann/drone-artifactory .

FROM ubuntu:17.04

ENV CLI_VERSION 1.12.0
RUN apt-get update && apt-get install -y \
  curl \
&& curl -L "https://jfrog.bintray.com/jfrog-cli-go/${CLI_VERSION}/jfrog-cli-linux-amd64/jfrog" -o /bin/jfrog \
&& apt-get remove -y curl \
&& apt-get autoremove -y \
&& apt-get clean \
&& rm -rf /var/lib/apt/lists/*

RUN chmod 0755 /bin/jfrog

ADD drone-artifactory /bin/
ENTRYPOINT ["/bin/drone-artifactory"]
