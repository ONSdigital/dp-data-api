#!/bin/bash -eux

pushd dp-data-api
  make build
  cp build/dp-data-api Dockerfile.concourse ../build
popd
