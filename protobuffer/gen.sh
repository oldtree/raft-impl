#!/bin/bash

echo $module
echo $PROTOC_INSTALL
echo $PWD
protoc -I$PROTOC_INSTALL/include -I. --go_out=plugins=grpc:. *.proto
cd 
echo `ls -h`