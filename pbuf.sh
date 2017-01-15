#!/bin/bash
cd "$(dirname "$0")"
cd proto
protoc --go_out=../web/src/pbuf ./*.proto
