run-user-server:
	@cd ./user && go run ./server/grpc/main.go

generate-proto:
	@sh ./proto-generator.sh