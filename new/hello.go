package main

import (
	"context"

	pb "github.com/forrestlinfeng/trpcdemo"
)

type helloImpl struct {
	pb.UnimplementedHello
}

// Hello Hello says hello.
func (s *helloImpl) Hello(
	ctx context.Context,
	req *pb.HelloRequest,
) (*pb.HelloResponse, error) {
	rsp := &pb.HelloResponse{}
	return rsp, nil
}
