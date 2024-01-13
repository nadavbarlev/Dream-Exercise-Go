#!/bin/bash

# `go build` is smart enough not to build the test files
# only if the build was succeeded then execute it. 
go build -o dream-exercise cmd/web/*.go && ./dream-exercise