package client

import (
	"context"

	pb "github.com/Questee29/taxi-app_orderService/proto/protob"
	"google.golang.org/grpc"
)

type Deps struct {
	ClientHandler pb.OrderGrpcServiceClient
}
type ClientConn struct {
	*grpc.ClientConn
	ctx    context.Context
	target string
	dopts  grpc.DialOption
}
type Client struct {
	deps Deps
	conn ClientConn
}

func NewClient(deps Deps, ctx context.Context, target string, dopts grpc.DialOption) *Client {
	return &Client{
		deps: deps,
		conn: ClientConn{
			ctx:    ctx,
			target: target,
			dopts:  dopts,
		},
	}
}

func (client *Client) Dial() (pb.OrderGrpcServiceClient, error) {
	//dial to the driverService
	conn, err := grpc.Dial(client.conn.target, client.conn.dopts) //
	if err != nil {
		return nil, err
	}
	c := pb.NewOrderGrpcServiceClient(conn)
	return c, nil
}
