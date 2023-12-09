package main

import (
	pb "github.com/forrestlinfeng/trpcdemo"
	_ "trpc.group/trpc-go/trpc-filter/debuglog"
	_ "trpc.group/trpc-go/trpc-filter/recovery"
	trpc "trpc.group/trpc-go/trpc-go"
	"trpc.group/trpc-go/trpc-go/log"
)

func main() {
	s := trpc.NewServer()
	pb.RegisterHelloService(s.Service("trpcdemo.Hello"), NewHelloImpl())
	if err := s.Serve(); err != nil {
		log.Fatal(err)
	}
}
