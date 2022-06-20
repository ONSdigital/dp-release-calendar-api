#!/bin/bash -eux

pushd dp-release-calendar-api
  make build
  cp build/dp-release-calendar-api Dockerfile.concourse ../build
popd
