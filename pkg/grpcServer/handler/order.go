package handler

import (
	"context"

	model "github.com/Questee29/taxi-app_orderService/models/order"
	pb "github.com/Questee29/taxi-app_orderService/proto/protob"
)

// var (
// 	business = make(chan int32, 50) // id int
// 	comfort  = make(chan int32, 50)
// 	economy  = make(chan int32, 50)
// )

const (
	DriverAdress = ":9081"
)

type OrderService interface {
	NewOrder(ctx context.Context, user model.UserRequest, driver model.DriverResponse) error
	FindFreeDriver(ctx context.Context, user model.UserRequest) (model.DriverResponse, error)
	//GetAllOrders(ctx context.Context, user model.UserRequest)
}

type OrderHandler struct {
	pb.UnimplementedOrderGrpcServiceServer
	service OrderService //GRPCOrderService
}

func NewOrderHandler(service OrderService) *OrderHandler {
	return &OrderHandler{
		service: service,
	}
}

func (h *OrderHandler) OrderTaxi(ctx context.Context, req *pb.OrderRequest) (*pb.OrderResponse, error) {

	// uReq := model.UserRequest{
	// 	ID:       req.Userid,
	// 	TaxiType: req.Type.String(),
	// 	From:     req.From,
	// 	To:       req.To,
	// }
	// go func() {
	// 	switch uReq.TaxiType {
	// 	case "comfort":
	// 		<-comfort
	// 	case "business":
	// 		<-business
	// 	case "economy":
	// 		<-economy
	// 	}
	// }()
	// select {
	// case <-comfort:
	// case <-business:
	// case <-economy:
	// case <-ctx.Done():
	// 	log.Println()
	// }
	// return &pb.OrderResponse{Driverid: 1, DriverName: "dsa", Type: 2, From: uReq.From, To: uReq.To}, nil
	return nil, nil
}

// func (h *OrderHandler) GetAllOrders(ctx context.Context, req *pb.GetOrdersRequest, opts ...grpc.CallOption) (*pb.GetOrdersResponse, error) {
// 	uReq := model.UserRequest{
// 		ID: req.Userid,
// 	}
// 	return &pb.GetOrdersResponse{Type: pb.CarType(uReq.ID), DriverName: "d", From: "moscow"}, nil

// }

// func toDriverPBModel(driver model.DriverResponse) *pb.OrderResponse {
// 	value, ok := pb.CarType_value[driver.TaxType]
// 	if !ok {
// 		log.Println("invalid car type")
// 	}
// 	return &pb.OrderResponse{
// 		Driverid: driver.Driverid,
// 		Type:     *pb.CarType(value).Enum(),
// 		From:     driver.From,
// 		To:       driver.To,
// 	}
// }
