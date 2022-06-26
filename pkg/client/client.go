package client

import (
	"log"

	pb "github.com/Questee29/taxi-app_orderService/proto/protob"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ClientOrder struct {
	orderClient pb.OrderGrpcServiceClient
}

func NewClientOrder(grpcConn grpc.ClientConnInterface) *ClientOrder {
	return &ClientOrder{
		orderClient: pb.NewOrderGrpcServiceClient(grpcConn),
	}
}
func New(target string) pb.OrderGrpcServiceClient {
	log.Println("dialling to order service....")
	conn, err := grpc.Dial(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("can not dial to the order service %s", err)
		return nil
	}
	log.Println("client successfully created!")
	return pb.NewOrderGrpcServiceClient(conn)

}
