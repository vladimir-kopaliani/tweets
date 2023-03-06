package templateservice

//go:generate protoc --go_out=./user_service/v1/ --go-grpc_out=:./user_service/v1/ ./user_service/v1/user_service.proto
