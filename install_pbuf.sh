#!/bin/sh
set -e
# check to see if protobuf folder is empty
if [ ! -d "$HOME/protobuf/lib" ]; then
  wget https://github.com/google/protobuf/releases/download/v3.1.0/protobuf-js-3.1.0.tar.gz;
  tar -xzf protobuf-js-3.1.0.tar.gz;
  cd protobuf-3.1.0 && ./configure --prefix=$HOME/protobuf && make && make install;
else
  echo 'Using cached directory.';
fi
