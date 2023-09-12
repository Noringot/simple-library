gen-docker:
	docker-compose up -d

gen-db:
	docker exec -i grpc-db-1 mysql -uroot -proot < ./db.sql

gen-proto:
	mkdir -p pkg/lib
	protoc \
	--go_out=pkg/lib --go_opt=paths=source_relative \
	--go-grpc_out=pkg/lib --go-grpc_opt=paths=source_relative \
	proto/lib/lib.proto
	mv pkg/lib/proto/lib/lib.pb.go pkg/lib
	mv pkg/lib/proto/lib/lib_grpc.pb.go pkg/lib
	rm -rd pkg/lib/proto

start-server:
	go run cmd/main/main.go
