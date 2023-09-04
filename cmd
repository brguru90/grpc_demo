export GO_PATH=~/go
export PATH=$PATH:/$GO_PATH/bin
protoc  --proto_path=../pb ../pb/*.proto --go_out=. --go-grpc_out=.