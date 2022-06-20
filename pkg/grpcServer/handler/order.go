package handler

import (
	"context"
	"log"

	model "github.com/Questee29/taxi-app_orderService/models/order"
	pb "github.com/Questee29/taxi-app_orderService/proto/protob"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	DriverAdress = ":9081"
)

type OrderService interface {
	NewOrder(ctx context.Context, user model.UserRequest, driver model.DriverResponse) error
	FindFreeDriver(ctx context.Context, user model.UserRequest) (model.DriverResponse, error)
}

type OrderHandler struct {
	pb.UnimplementedOrderGrpcServer
	service OrderService //GRPCOrderService
}

func NewOrderHandler(service OrderService) *OrderHandler {
	return &OrderHandler{
		service: service,
	}
}

func (h *OrderHandler) OrderTaxi(ctx context.Context, req *pb.OrderRequest) (*pb.OrderResponse, error) {
	uReq := model.UserRequest{
		ID:       req.Userid,
		TaxiType: req.Type.String(),
		From:     req.From,
		To:       req.To,
	}
	log.Println(uReq)
	//dial to the driverService
	conn, err := grpc.Dial(DriverAdress, grpc.WithTransportCredentials(insecure.NewCredentials())) //
	if err != nil {

		return nil, err
	}
	defer conn.Close()

	c := pb.NewOrderGrpcClient(conn)
	result, err := c.FindDriver(ctx, &pb.FindDriverRequest{Userid: uReq.ID, Type: req.GetType()})
	if err != nil {
		return nil, err
	}
	//create model
	dResponse := model.DriverResponse{
		Driverid: result.GetDriverid(),
		TaxType:  result.GetType().String(),
		From:     result.GetFrom(),
		To:       result.GetTo(),
	}
	//create new order, if returned driver
	if err := h.service.NewOrder(ctx, uReq, dResponse); err != nil {
		log.Println("ERROR!")
		return nil, err
	}

	//returns driver
	return toDriverPBModel(dResponse), nil

}

func toDriverPBModel(driver model.DriverResponse) *pb.OrderResponse {
	value, ok := pb.CarType_value[driver.TaxType]
	if !ok {
		log.Println("invalid car type")
	}
	return &pb.OrderResponse{
		Driverid: driver.Driverid,
		Type:     *pb.CarType(value).Enum(),
		From:     driver.From,
		To:       driver.To,
	}
}
