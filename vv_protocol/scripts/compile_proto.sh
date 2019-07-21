#!/usr/bin sh

PROJECT_ROOT=${PROJECT_ROOT}
if [[ ${PROJECT_ROOT} == "" ]]; then
    echo '${PROJECT_ROOT} is required.'
    exit;
fi

PROTOBUF_DIR=${PROJECT_ROOT}/proto
function compile_golang() {
  GO_OUTPUT=${PROJECT_ROOT}/go/protocol
  ls ${PROTOBUF_DIR} | xargs protoc --proto_path=${PROTOBUF_DIR} --go_out=plugins=grpc:${GO_OUTPUT}
}

function compile_javascript() {
  JS_OUTPUT=${PROJECT_ROOT}/js/protocol
  ls ${PROTOBUF_DIR} | xargs protoc --proto_path=${PROTOBUF_DIR} \
                                    --js_out=import_style=commonjs:${JS_OUTPUT} \
                                    --grpc-web_out=import_style=typescript,mode=grpcwebtext:${JS_OUTPUT}
}

case "$1" in
  "golang"     ) compile_golang;;
  "js" ) compile_javascript;;
  *            ) echo "unexpected language or no args given.";;
esac

exit 0