#!/bin/bash
cd "$(dirname "$0")"
cd proto
protoc --go_out=../web/src/memmi/pbuf ./memmi.proto
protoc --go_out=../web/src/memmi/config ./config.proto
pbjs -t static-module -w es6 ./memmi.proto > ../client-side/src/pbuf/pbuf.js

