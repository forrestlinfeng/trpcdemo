package main

import (
	"context"

	pb "github.com/forrestlinfeng/trpcdemo"
	"trpc.group/trpc-go/trpc-go"
	"trpc.group/trpc-go/trpc-go/log"

	"github.com/polarismesh/polaris-go"
	"github.com/polarismesh/polaris-go/pkg/model"
)

type helloImpl struct {
	configFile polaris.ConfigFile
	val        string
}

func NewHelloImpl() pb.HelloService {
	configAPI, err := polaris.NewConfigAPIByFile(trpc.ServerConfigPath)
	// configAPI, err := polaris.NewConfigAPI()
	if err != nil {
		log.Errorf("fail new api: %v", err)
	}

	configFile, err := configAPI.GetConfigFile("Polaris", "test", "aa.json")
	if err != nil {
		log.Errorf("fail to get config.", err)
	}
	s := &helloImpl{}
	configFile.AddChangeListener(s.changeListener)
	s.configFile = configFile
	s.val = configFile.GetContent()
	log.Infof("content: %v", s.val)
	return s
}

// Hello Hello says hello.
func (s *helloImpl) Hello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Msg: s.val}, nil
}

func (s *helloImpl) changeListener(event model.ConfigFileChangeEvent) {
	log.Infof("received change event. %+v, value:%v", event, event.NewValue)
	s.val = event.NewValue
}
