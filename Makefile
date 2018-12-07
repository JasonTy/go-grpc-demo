GO=go run

run-client:
	${GO} client/client.go

run-server:
	${GO} server/server.go

run-protoc:
	protoc --go_out=plugins=grpc:. protocol/helloword.proto
