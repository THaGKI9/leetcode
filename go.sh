#/bin/bash

cd main
go run $(ls -1 *.go | grep -v _test.go) -- $@
