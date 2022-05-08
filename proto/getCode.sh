ROOT=$(dirname "${BASH_SOURCE[0]}")
echo $ROOT
cd $ROOT && pwd

protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative wgSvc.proto 