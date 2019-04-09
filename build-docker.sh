#!/bin/bash

DOCKER_IMAGE="jmccann/drone-artifactory"

tag=$1

if [ -z $tag ]; then
  echo "please provide a tag arg"
  exit 1
fi

major=$(echo $tag | awk -F. '{print $1}')
minor=$(echo $tag | awk -F. '{print $2}')

echo "Confirm building images for:"
echo "  MAJOR: ${major}"
echo "  MINOR: ${minor}"

read -p "Proceed? [Y/N] " ans

if [[ "$ans" != "Y" && "$ans" != "y" ]]; then
  echo "Cancelling"
  exit 0
fi

set -x
docker build -t ${DOCKER_IMAGE}:latest .

docker tag ${DOCKER_IMAGE}:latest ${DOCKER_IMAGE}:${major}
docker tag ${DOCKER_IMAGE}:latest ${DOCKER_IMAGE}:${major}.${minor}

exit 0

docker push ${DOCKER_IMAGE}:latest
docker push ${DOCKER_IMAGE}:${major}
docker push ${DOCKER_IMAGE}:${major}.${minor}
set +x
