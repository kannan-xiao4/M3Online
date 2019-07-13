#sh

REPO_ROOT=$(git rev-parse --show-toplevel)
SERVER_SRC=$REPO_ROOT/server/src/m3online/
SERVER_OUTPUT_DIR=$SERVER_SRC/rpc/
CLIENT_OUTPUT_DIR=$REPO_ROOT/unityclient/Assets/M3Online/Grpc/
GRPC_SOURCE=$REPO_ROOT/protocolbuffer

mkdir -p $CLIENT_OUTPUT_DIR
mkdir -p $SERVER_OUTPUT_DIR
rm -rf $CLIENT_OUTPUT_DIR/*.cs
rm -rf $SERVER_OUTPUT_DIR/*.go

protoc --proto_path=$GRPC_SOURCE \
    --csharp_out=$CLIENT_OUTPUT_DIR/ \
    --go_out=$SERVER_OUTPUT_DIR \
    --grpc_out=$CLIENT_OUTPUT_DIR/ \
    -I=$GRPC_SOURCE/ \
    $GRPC_SOURCE/*/*.proto  \
    --plugin=protoc-gen-grpc=/usr/local/bin/grpc_csharp_plugin
