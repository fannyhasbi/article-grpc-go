run-user-server:
	@cd ./user && go run ./server/grpc/main.go

run-article-client:
	@cd ./article && go run ./client/grpc/main.go

generate-proto:
	@sh ./proto-generator.sh