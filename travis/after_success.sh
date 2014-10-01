#!/bin/bash
set -e
go get code.google.com/p/go.tools/cmd/cover
go get github.com/modocache/gover
go get github.com/mattn/goveralls

cd go
go list -f '{{if len .TestGoFiles}}go test -coverprofile={{.Dir}}/.coverprofile {{.ImportPath}}{{end}}' ./... | xargs -i sh -c {}
gover
goveralls -coverprofile=gover.coverprofile -repotoken $COVERALLS_TOKEN
cd ..
