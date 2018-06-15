#!/bin/bash

echo $module
echo $PROTOC_INSTALL
echo $PWD
protoc -I$PROTOC_INSTALL/include -I. --go_out=plugins=grpc:. raft.proto
protoc -I$PROTOC_INSTALL/include -I. --go_out=plugins=grpc:. extern-api.proto
cd 
echo `ls -h`