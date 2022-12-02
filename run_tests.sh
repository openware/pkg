#!/bin/sh
set -xe

if [ -z ${COVER_OUT} ]; then
  echo "COVER_OUT is required"
  exit 1
fi

ROOT_DIR=$PWD

for d in */ ; do
	cd $d

	go mod download
	go test ./... --cover -coverprofile $COVER_OUT
	go vet ./...

	cd $ROOT_DIR
done
