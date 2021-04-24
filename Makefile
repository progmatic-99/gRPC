.PHONY:	protos
protos:
	protoc -I protos/ ./currency.proto --go-grpc_out=currency/