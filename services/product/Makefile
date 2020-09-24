.PHONY: proto

proto:
	protoc ../../protos/*.proto --proto_path=../../protos/ --go_out=plugins=grpc:${PWD}/proto 
run:
	DATABASE_URL=postgresql://postgres:admin@0.0.0.0:5432/ecommerce?sslmode=disable DISCOUNT_ADDR=0.0.0.0:3003 go run main.go
mocks:
	mockery -all