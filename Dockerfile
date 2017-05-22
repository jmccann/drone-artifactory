# Docker image for the Drone Artifactory plugin
#
#     docker build -t jmccann/drone-artifactory .

FROM ubuntu:17.04

RUN apt-get update && apt-get install -y \
  curl \
&& curl -fL https://getcli.jfrog.io | sh \
&& apt-get remove -y curl \
&& apt-get autoremove -y \
&& apt-get clean \
&& rm -rf /var/lib/apt/lists/*

RUN mv jfrog /bin/jfrog

ADD drone-artifactory /bin/
ENTRYPOINT ["/bin/drone-artifactory"]
