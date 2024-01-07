#!/usr/bin/env bash

export GOGC=off

go build -gcflags=-B .

time ./go1brc "$1"