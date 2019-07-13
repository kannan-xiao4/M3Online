# C#
protoc -I ./protocolbuffer \
    --csharp_out unity/Assets/M3Online/Protoc \
    --grpc_out unity/Assets/M3Online/Protoc ./protocolbuffer/*.proto \
    --plugin=protoc-gen-grpc=/usr/local/bin/grpc_csharp_plugin
 
# Go
protoc -I ./protocolbuffer --go_out=plugins=grpc:./server/src/m3online ./protocolbuffer/*.proto