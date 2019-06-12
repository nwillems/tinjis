#!/bin/bash

# You could probably do something smarter with travis. meh
VERSION="$(git rev-parse HEAD)"

docker build -t nwillems/payment:${VERSION} .

if [ "${DEPLOY}" == "true" ]; then
    docker push nwillems/payment:${VERSION}
fi
