package server

import (
	"log"
	"golang-gRPC-example/api"
	"golang.org/x/net/context"
)

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, req *api.Request) (*api.Response, error) {
	log.Printf("{Id:%d, Username:%s, Password:%s, Active: %t, GPA: %F", req.Id, req.Username, req.Password, req.Active, req.Gpa)
	return &api.Response{Status: 1, Message: "The data was successfully accepted!!!"}, nil
}
