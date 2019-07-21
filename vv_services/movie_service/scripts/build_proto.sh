#!/usr/bin/env bash

PROJECT_ROOT=${PROJECT_ROOT}
if [[ ${PROJECT_ROOT} == "" ]]; then
    echo '${PROJECT_ROOT} is required.'
    exit;
fi

PROTOBUF_DIR=${PROJECT_ROOT}/internal/app/interface/rpc/v1/protocol
PROTOBUF_FILE=${PROTOBUF_DIR}/movie.proto

#MOVIE_SERVICE=${PROJECT_ROOT}/internal/app/interface/rpc/v1

# golang
protoc --proto_path=${PROTOBUF_DIR} ${PROTOBUF_FILE} --go_out=plugins=grpc:${PROTOBUF_DIR}