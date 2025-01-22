run:
		go run cmd/application/main.go

gen-rpc-files:
		protoc \
		--go_out=. \
		--go-grpc_out=. \
		./internal/api/grpc/pb/*.proto
