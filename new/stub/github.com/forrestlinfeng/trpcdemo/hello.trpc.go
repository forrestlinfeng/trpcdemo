// Code generated by trpc-go/trpc-cmdline v1.0.5. DO NOT EDIT.
// source: hello.proto

package trpcdemo

import (
	"context"
	"errors"
	"fmt"

	_ "trpc.group/trpc-go/trpc-go"
	"trpc.group/trpc-go/trpc-go/client"
	"trpc.group/trpc-go/trpc-go/codec"
	_ "trpc.group/trpc-go/trpc-go/http"
	"trpc.group/trpc-go/trpc-go/server"
)

// START ======================================= Server Service Definition ======================================= START

// HelloService defines service.
type HelloService interface {
	// Hello Hello says hello.
	Hello(ctx context.Context, req *HelloRequest) (*HelloResponse, error)
}

func HelloService_Hello_Handler(svr interface{}, ctx context.Context, f server.FilterFunc) (interface{}, error) {
	req := &HelloRequest{}
	filters, err := f(req)
	if err != nil {
		return nil, err
	}
	handleFunc := func(ctx context.Context, reqbody interface{}) (interface{}, error) {
		return svr.(HelloService).Hello(ctx, reqbody.(*HelloRequest))
	}

	var rsp interface{}
	rsp, err = filters.Filter(ctx, req, handleFunc)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

// HelloServer_ServiceDesc descriptor for server.RegisterService.
var HelloServer_ServiceDesc = server.ServiceDesc{
	ServiceName: "trpcdemo.Hello",
	HandlerType: ((*HelloService)(nil)),
	Methods: []server.Method{
		{
			Name: "/trpcdemo.Hello/Hello",
			Func: HelloService_Hello_Handler,
		},
	},
}

// RegisterHelloService registers service.
func RegisterHelloService(s server.Service, svr HelloService) {
	if err := s.Register(&HelloServer_ServiceDesc, svr); err != nil {
		panic(fmt.Sprintf("Hello register error:%v", err))
	}
}

// START --------------------------------- Default Unimplemented Server Service --------------------------------- START

type UnimplementedHello struct{}

// Hello Hello says hello.
func (s *UnimplementedHello) Hello(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {
	return nil, errors.New("rpc Hello of service Hello is not implemented")
}

// END --------------------------------- Default Unimplemented Server Service --------------------------------- END

// END ======================================= Server Service Definition ======================================= END

// START ======================================= Client Service Definition ======================================= START

// HelloClientProxy defines service client proxy
type HelloClientProxy interface {
	// Hello Hello says hello.
	Hello(ctx context.Context, req *HelloRequest, opts ...client.Option) (rsp *HelloResponse, err error)
}

type HelloClientProxyImpl struct {
	client client.Client
	opts   []client.Option
}

var NewHelloClientProxy = func(opts ...client.Option) HelloClientProxy {
	return &HelloClientProxyImpl{client: client.DefaultClient, opts: opts}
}

func (c *HelloClientProxyImpl) Hello(ctx context.Context, req *HelloRequest, opts ...client.Option) (*HelloResponse, error) {
	ctx, msg := codec.WithCloneMessage(ctx)
	defer codec.PutBackMessage(msg)
	msg.WithClientRPCName("/trpcdemo.Hello/Hello")
	msg.WithCalleeServiceName(HelloServer_ServiceDesc.ServiceName)
	msg.WithCalleeApp("")
	msg.WithCalleeServer("")
	msg.WithCalleeService("Hello")
	msg.WithCalleeMethod("Hello")
	msg.WithSerializationType(codec.SerializationTypePB)
	callopts := make([]client.Option, 0, len(c.opts)+len(opts))
	callopts = append(callopts, c.opts...)
	callopts = append(callopts, opts...)
	rsp := &HelloResponse{}
	if err := c.client.Invoke(ctx, req, rsp, callopts...); err != nil {
		return nil, err
	}
	return rsp, nil
}

// END ======================================= Client Service Definition ======================================= END
