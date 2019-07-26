# user service proto
protoc --go_out=plugins=grpc:. ./user/proto/*.proto

# article service proto
protoc --go_out=plugins=grpc:. ./article/proto/*.proto