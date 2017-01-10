all: grpcexample/person.pb.go client server

grpcexample/person.pb.go:
	protoc grpcexample/person.proto --go_out=plugins=grpc:.

client:
	go build -tags=client -o client

server:
	go build -tags=server -o server 

clean:
	rm grpcexample/person.pb.go
	rm client
	rm server

test:
	go test -v -tags=server -benchmem -bench .
