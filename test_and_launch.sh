#!/bin/bash
go test ./...
[ $? -eq 0 ] || exit $?;
echo continue
go run poemserver/main.go './fake/path'
