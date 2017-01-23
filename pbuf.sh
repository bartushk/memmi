#!/bin/bash
cd "$(dirname "$0")"
cd proto
protoc --go_out=../web/src/memmi/pbuf ./*.proto
pbjs -t static-module -w es6 ./*.proto |\
sed 's/!== ""/=== null ? "" : true/' |\
sed 's/!== 0/=== null ? 0 : true/' |\
sed 's/!== false/=== false ? true : true/' >\
../client-side/src/pbuf/pbuf.js
