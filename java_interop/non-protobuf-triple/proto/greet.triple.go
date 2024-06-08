// Code generated by protoc-gen-triple. DO NOT EDIT.
//
// Source: greet.proto
package greet

import (
	"context"
)

import (
	"dubbo.apache.org/dubbo-go/v3"
	"dubbo.apache.org/dubbo-go/v3/client"
	"dubbo.apache.org/dubbo-go/v3/common"
	"dubbo.apache.org/dubbo-go/v3/common/constant"
	"dubbo.apache.org/dubbo-go/v3/protocol/triple/triple_protocol"
	"dubbo.apache.org/dubbo-go/v3/server"
)

// This is a compile-time assertion to ensure that this generated file and the Triple package
// are compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of Triple newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of Triple or updating the Triple
// version compiled into your binary.
const _ = triple_protocol.IsAtLeastVersion0_1_0

const (
	// GreetingsServiceName is the fully-qualified name of the GreetingsService service.
	GreetingsServiceName = "org.apache.dubbo.tri.hessian2.api.GreetingsService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// GreetingsServiceGreetProcedure is the fully-qualified name of the GreetingsService's Greet RPC.
	GreetingsServiceGreetProcedure = "/org.apache.dubbo.tri.hessian2.api.GreetingsService/Greet"
)

var (
	_ GreetingsService = (*GreetingsServiceImpl)(nil)
)

// GreetingsService is a client for the org.apache.dubbo.tri.hessian2.api.GreetingsService service.
type GreetingsService interface {
	Greet(ctx context.Context, req *GreetRequest, opts ...client.CallOption) (*GreetResponse, error)
}

// NewGreetingsService constructs a client for the greet.GreetingsService service.
func NewGreetingsService(cli *client.Client, opts ...client.ReferenceOption) (GreetingsService, error) {
	conn, err := cli.DialWithInfo("org.apache.dubbo.tri.hessian2.api.GreetingsService", &GreetingsService_ClientInfo, opts...)
	if err != nil {
		return nil, err
	}
	return &GreetingsServiceImpl{
		conn: conn,
	}, nil
}

func SetConsumerGreetingsService(srv common.RPCService) {
	dubbo.SetConsumerServiceWithInfo(srv, &GreetingsService_ClientInfo)
}

// GreetingsServiceImpl implements GreetingsService.
type GreetingsServiceImpl struct {
	conn *client.Connection
}

func (c *GreetingsServiceImpl) Greet(ctx context.Context, req *GreetRequest, opts ...client.CallOption) (*GreetResponse, error) {
	resp := new(GreetResponse)
	if err := c.conn.CallUnary(ctx, []interface{}{req}, resp, "Greet", opts...); err != nil {
		return nil, err
	}
	return resp, nil
}

var GreetingsService_ClientInfo = client.ClientInfo{
	InterfaceName: "org.apache.dubbo.tri.hessian2.api.GreetingsService",
	MethodNames:   []string{"Greet"},
	ConnectionInjectFunc: func(dubboCliRaw interface{}, conn *client.Connection) {
		dubboCli := dubboCliRaw.(*GreetingsServiceImpl)
		dubboCli.conn = conn
	},
}

// GreetingsServiceHandler is an implementation of the org.apache.dubbo.tri.hessian2.api.GreetingsService service.
type GreetingsServiceHandler interface {
	Greet(context.Context, *GreetRequest) (*GreetResponse, error)
}

func RegisterGreetingsServiceHandler(srv *server.Server, hdlr GreetingsServiceHandler, opts ...server.ServiceOption) error {
	return srv.Register(hdlr, &GreetingsService_ServiceInfo, opts...)
}

func SetProviderGreetingsService(srv common.RPCService) {
	dubbo.SetProviderServiceWithInfo(srv, &GreetingsService_ServiceInfo)
}

var GreetingsService_ServiceInfo = server.ServiceInfo{
	InterfaceName: "org.apache.dubbo.tri.hessian2.api.GreetingsService",
	ServiceType:   (*GreetingsServiceHandler)(nil),
	Methods: []server.MethodInfo{
		{
			Name: "Greet",
			Type: constant.CallUnary,
			ReqInitFunc: func() interface{} {
				return new(GreetRequest)
			},
			MethodFunc: func(ctx context.Context, args []interface{}, handler interface{}) (interface{}, error) {
				req := args[0].(*GreetRequest)
				res, err := handler.(GreetingsServiceHandler).Greet(ctx, req)
				if err != nil {
					return nil, err
				}
				return triple_protocol.NewResponse(res), nil
			},
		},
	},
}
