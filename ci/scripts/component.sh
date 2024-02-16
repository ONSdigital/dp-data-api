#!/bin/bash -eux

pushd dp-data-api
  make test-component
popd
