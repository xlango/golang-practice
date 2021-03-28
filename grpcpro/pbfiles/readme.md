
    protoc --go_out=plugins=grpc:../services/ Prod.proto
    protoc --grpc-gateway_out=logtostderr=true:../services/ Prod.proto