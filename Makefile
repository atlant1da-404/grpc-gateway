
# Step one .proto -> pb.go !
create:
	protoc -I ./proto \
          --go_out ./proto --go_opt paths=source_relative \
          --go-grpc_out ./proto --go-grpc_opt paths=source_relative \
          --grpc-gateway_out ./proto --grpc-gateway_opt paths=source_relative \
          proto/notes/notes.proto

run:
	go run server/server.go

# Run redis in docker
redis:
	docker run --name rdb -p 6379:6379 redis
