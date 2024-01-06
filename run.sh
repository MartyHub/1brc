#!/usr/bin/env bash

go build .

export GOGC=off

time ./go1brc "$1"