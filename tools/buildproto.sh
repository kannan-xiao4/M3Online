#sh

REPO_ROOT=$(git rev-parse --show-toplevel)
SERVER_SRC=$REPO_ROOT/server/src/m3online/
SERVER_OUTPUT_DIR=$SERVER_SRC/rpc/
CLIENT_OUTPUT_DIR=$REPO_ROOT/unityclient/Assets/M3Online/Grpc/
GRPC_SOURCE=$REPO_ROOT/protocolbuffer
PROTO_FILES=$GRPC_SOURCE/*/*.proto

mkdir -p $CLIENT_OUTPUT_DIR
mkdir -p $SERVER_OUTPUT_DIR
rm -rf $CLIENT_OUTPUT_DIR/*.cs
rm -rf $SERVER_OUTPUT_DIR/*.go

#golang
protoc -I=$GRPC_SOURCE \
    --go_out=plugins=grpc:$SERVER_OUTPUT_DIR $PROTO_FILES

#C#
protoc -I=$GRPC_SOURCE \
    --csharp_out=$CLIENT_OUTPUT_DIR \
    --grpc_out=$CLIENT_OUTPUT_DIR $PROTO_FILES \
    --plugin=protoc-gen-grpc=/usr/local/bin/grpc_csharp_plugin    
