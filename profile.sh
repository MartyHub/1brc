#!/usr/bin/env bash

export GOGC=off

go test -cpuprofile cpu.prof
pprof -http=:8080 cpu.prof
