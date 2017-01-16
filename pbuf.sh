#!/bin/bash
cd "$(dirname "$0")"
cd proto
protoc --go_out=../web/src/pbuf ./*.proto
pbjs -t static-module -w es6 ./*.proto > ../client-side/src/pbuf/pbuf.js
