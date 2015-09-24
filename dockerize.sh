#!/bin/bash
set -o xtrace
echo "start dockerizing ..."
export FILE='./Makefile'
if [ -f $FILE ]
then
       echo "File $FILE exist."
else
       echo "File $FILE does not exist."
       go generate
fi
echo "start dockerizing ..."
go build
godep save -r
docker build -t mingderwang/onion .
docker run -i -t -v $(pwd):/gopath/src/onion mingderwang/onion onion migratedb
docker run -d -p 8080:8080 -v $(pwd):/gopath/src/onion mingderwang/onion onion serve
