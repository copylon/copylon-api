package internal

import (
	"google.golang.org/grpc"
)

func Multiplex() *grpc.Server {
	s := grpc.NewServer(grpc.StreamInterceptor(nil), grpc.UnaryInterceptor(nil))
	//s.Serve(":5001")
	return s
}
