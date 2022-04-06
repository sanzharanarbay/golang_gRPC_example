package main


import (
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"golang-gRPC-example/api"
)

func main() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000",  grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := api.NewChatServiceClient(conn)

	student := api.Request{
		Id: 15,
		Username: "sanzhar",
		Password: "qwerty123!",
		Active: true,
		Gpa: 3.85,
	}

		response, err := c.SayHello(context.Background(), &student)
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("{Status:%d, Message:%s}", response.Status, response.Message)

}