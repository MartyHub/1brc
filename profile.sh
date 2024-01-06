#!/usr/bin/env bash

go build .

export GOGC=off

go test -cpuprofile cpu.prof -bench .
pprof -http=:8080 cpu.prof
