#!/usr/bin/bash

set -x

rm -rf ./bin
rm -rf ./log
rm -rf ./output
mkdir bin
mkdir log
mkdir output

go mod tidy
go build -o ./bin/mini_spider
cd bin
./mini_spider -c ../conf -l ../log