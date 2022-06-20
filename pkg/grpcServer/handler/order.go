package handler

import (
	"context"
	"fmt"
	"log"

	model "github.com/Questee29/taxi-app_orderService/models/order"
	pb "github.com/Questee29/taxi-app_orderService/proto/protob"
)

type OrderService interface {
	OrderTaxi(ctx context.Context, user model.UserRequest) error
}

type OrderHandler struct {
	pb.UnimplementedOrderGrpcServer
	service OrderService
}

func NewOrderHandler(service OrderService) *OrderHandler {
	return &OrderHandler{
		service: service,
	}
}

func (h *OrderHandler) OrderTaxi(ctx context.Context, req *pb.OrderRequest) (*pb.OrderResponse, error) {
	u := model.UserRequest{
		ID:       req.Userid,
		TaxiType: req.Type.String(),
		From:     req.From,
		To:       req.To,
	}
	fmt.Println(u.TaxiType)
	if err := h.service.OrderTaxi(ctx, u); err != nil {
		return nil, err
	}
	return toPBModel(u), nil

}
func toPBModel(user model.UserRequest) *pb.OrderResponse {
	value, ok := pb.CarType_value[user.TaxiType]
	if !ok {
		log.Println("invalid car type")
	}
	return &pb.OrderResponse{
		Driverid: user.ID,
		Type:     *pb.CarType(value).Enum(),
		From:     user.From,
		To:       user.To,
	}
}
