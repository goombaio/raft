#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

if [ ! $(command -v gometalinter) ]
then
	go get github.com/alecthomas/gometalinter
	gometalinter --update --install
fi

time gometalinter \
	--exclude='error return value not checked.*(Close|Log|Print).*\(errcheck\)$' \
	--exclude='.*_test\.go:.*error return value not checked.*\(errcheck\)$' \
	--exclude='/thrift/' \
	--exclude='/pb/' \
	--exclude='no args in Log call \(vet\)' \
	--disable=dupl \
	--disable=aligncheck \
	--disable=gotype \
	--cyclo-over=20 \
	--tests \
	--concurrency=2 \
	--deadline=300s \
	./...