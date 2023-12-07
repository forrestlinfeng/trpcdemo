package main

import (
	"context"
	"reflect"
	"testing"

	pb "github.com/forrestlinfeng/trpcdemo"
	"github.com/golang/mock/gomock"
	_ "trpc.group/trpc-go/trpc-go/http"
)

//go:generate go mod tidy
//go:generate mockgen -destination=stub/github.com/forrestlinfeng/trpcdemo/hello_mock.go -package=trpcdemo -self_package=github.com/forrestlinfeng/trpcdemo --source=stub/github.com/forrestlinfeng/trpcdemo/hello.trpc.go

func Test_helloImpl_Hello(t *testing.T) {
	// Start writing mock logic.
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	helloService := pb.NewMockHelloService(ctrl)
	var inorderClient []*gomock.Call
	// Expected behavior.
	m := helloService.EXPECT().Hello(gomock.Any(), gomock.Any()).AnyTimes()
	m.DoAndReturn(func(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
		s := &helloImpl{}
		return s.Hello(ctx, req)
	})
	gomock.InOrder(inorderClient...)

	// Start writing unit test logic.
	type args struct {
		ctx context.Context
		req *pb.HelloRequest
		rsp *pb.HelloResponse
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var rsp *pb.HelloResponse
			var err error
			if rsp, err = helloService.Hello(tt.args.ctx, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("helloImpl.Hello() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(rsp, tt.args.rsp) {
				t.Errorf("helloImpl.Hello() rsp got = %v, want %v", rsp, tt.args.rsp)
			}
		})
	}
}
