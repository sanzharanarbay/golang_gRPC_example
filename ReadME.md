# Goland gRPC client and server Example

- go get -u github.com/golang/protobuf/protoc-gen-go;
- go get golang.org/x/net/context@v0.0.0-20200822124328-c89045814202;
- protoc proto/chat.proto  --go_out=plugins=grpc:api proto/chat.proto ;
-  Simular Commands : 
* protoc -I proto --go_out=plugins=grpc:chat proto/chat.proto;
* protoc --go_out=plugins=grpc:chat chat.proto;
  
-  In Ubuntu if error about protobuf, then run sudo apt  install protobuf-compiler; 
-  Run Server: go run main.go ;
- Run Client: go run client.go ;

